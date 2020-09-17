#include <math.h>
#include <stdbool.h>
#include <stdio.h>

/* If the eccentricity is very close to parabolic,  and the eccentric
anomaly is quite low,  you can get an unfortunate situation where
roundoff error keeps you from converging.  Consider the just-barely-
elliptical case,  where in Kepler's equation,

M = E - e sin( E)

   E and e sin( E) can be almost identical quantities.  To
around this,  near_parabolic( ) computes E - e sin( E) by expanding
the sine function as a power series:

E - e sin( E) = E - e( E - E^3/3! + E^5/5! - ...)
= (1-e)E + e( -E^3/3! + E^5/5! - ...)

   It's a little bit expensive to do this,  and you only need do it
quite rarely.  (I only encountered the problem because I had orbits
that were supposed to be 'pure parabolic',  but due to roundoff,
they had e = 1+/- epsilon,  with epsilon _very_ small.)  So 'near_parabolic'
is only called if we've gone seven iterations without converging. */

static double near_parabolic( const double ecc_anom, const double e)
{
   const double anom2 = (e > 1. ? ecc_anom * ecc_anom : -ecc_anom * ecc_anom);
   double term = e * anom2 * ecc_anom / 6.;
   double rval = (1. - e) * ecc_anom - term;
   unsigned n = 4;

   while( fabs( term) > 1e-15)
      {
      term *= anom2 / (double)(n * (n + 1));
      rval -= term;
      n += 2;
      }
   return( rval);
}

#define MAX_ITERATIONS 7
#define THRESH 1.e-12
#define MIN_THRESH 1.e-15
#define CUBE_ROOT( X)  (exp( log( X) / 3.))
#define PI 3.1415926535897932384626433832795028841971693993751058209749445923

static double kepler( const double ecc, double mean_anom)
{
   double curr, err, thresh, offset = 0.;
   double delta_curr = 1.;
   bool is_negative = false;
   unsigned n_iter = 0;

   if( !mean_anom)
      return( 0.);

   if( ecc < 1.)
      {
      if( mean_anom < -PI || mean_anom > PI)
         {
         double tmod = fmod( mean_anom, PI * 2.);

         if( tmod > PI)             /* bring mean anom within -pi to +pi */
            tmod -= 2. * PI;
         else if( tmod < -PI)
            tmod += 2. * PI;
         offset = mean_anom - tmod;
         mean_anom = tmod;
         }

      if( ecc < .99999)     /* low-eccentricity formula from Meeus,  p. 195 */
         {
         curr = atan2( sin( mean_anom), cos( mean_anom) - ecc);
         do
            {
            err = (curr - ecc * sin( curr) - mean_anom) / (1. - ecc * cos( curr));
            curr -= err;
            }
            while( fabs( err) > THRESH);
         return( curr + offset);
         }
      }


   if( mean_anom < 0.)
      {
      mean_anom = -mean_anom;
      is_negative = true;
      }

   curr = mean_anom;
   thresh = THRESH * fabs( 1. - ecc);
               /* Due to roundoff error,  there's no way we can hope to */
               /* get below a certain minimum threshhold anyway:        */
   if( thresh < MIN_THRESH)
      thresh = MIN_THRESH;
   if( thresh > THRESH)       /* i.e.,  ecc > 2. */
      thresh = THRESH;
   if( mean_anom < PI / 3. || ecc > 1.)    /* up to 60 degrees */
      {
      double trial = mean_anom / fabs( 1. - ecc);

      if( trial * trial > 6. * fabs(1. - ecc))   /* cubic term is dominant */
         {
         if( mean_anom < PI)
            trial = CUBE_ROOT( 6. * mean_anom);
         else        /* hyperbolic w/ 5th & higher-order terms predominant */
            trial = asinh( mean_anom / ecc);
         }
      curr = trial;
      }
   if( ecc < 1.)
      while( fabs( delta_curr) > thresh)
         {
         if( n_iter++ > MAX_ITERATIONS)
            err = near_parabolic( curr, ecc) - mean_anom;
         else
            err = curr - ecc * sin( curr) - mean_anom;
         delta_curr = -err / (1. - ecc * cos( curr));
         curr += delta_curr;
         }
   else
      while( fabs( delta_curr) > thresh)
         {
         if( n_iter++ > MAX_ITERATIONS)
            err = -near_parabolic( curr, ecc) - mean_anom;
         else
            err = ecc * sinh( curr) - curr - mean_anom;
         delta_curr = -err / (ecc * cosh( curr) - 1.);
         curr += delta_curr;
         }
   return( is_negative ? offset - curr : offset + curr);
}



int main()
{
   //9254
   //
         double _kepler;
        _kepler = kepler(5791,304.7431);
         
         printf("kepler value: %lf \n",_kepler);


}

/*
{
 "name": "ENDUROSAT ONE",
 "line1": {
  "linenumber": 1,
  "satellitenumber": 43551,
  "elsetclassification": "U",
  "internationaldesignator": "98067NZ",
  "elementsetepoch": "20246.3718749",
  "timemotionderivativeone": " .00145727",
  "timemotionderivativetwo": " 19081-4",
  "bdragterm": " 43614-3",
  "elementsetype": 0,
  "elementnumber": 999,
  "checksum": 5
 },
 "line2": {
  "linenumber": 2,
  "satellitenumber": 43551,
  "elseclassification": "",
  "orbitinclination": 51.6278,
  "raan": 236.8595,
  "eccentricity": 5791,
  "argumentperigee": 55.4121,
  "meananomaly": 304.7431,
  "meanmotion": 15.92046635,
  "revolutionnumber": 2258,
  "checksum": 8
 }
}
*/