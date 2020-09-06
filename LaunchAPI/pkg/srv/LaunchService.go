package srv

/*
function CalculateTargetVelocity()
	{
		var Starting_point_Altitude = parseFloat(elStarting_point_Altitude.value); --> Space Port
		var Additional_Speed = parseFloat(elAdditional_Speed.value); --> Space Port
		var Orbit_Perigee = parseFloat(elOrbit_Perigee.value); --> Orbit
		var Orbit_Apogee = parseFloat(elOrbit_Apogee.value); --> Orbit
		var Spaceport_latitude = parseFloat(elSpaceport_latitude.value); --> Space Port
		var Orbit_Inclination; -- > orbit (NOT SET) --> Orbit
		if (elOrbit.selectedIndex == 2) -- > Orbit == SSO
		{
			Orbit_Inclination = 95.7+.00296*Orbit_Perigee+7.97e-7*Orbit_Perigee*Orbit_Perigee;
			elOrbit_Inclination.value = Orbit_Inclination.toFixed(2);
		} else
			Orbit_Inclination = parseFloat(elOrbit_Inclination.value); -- > Default unless SSO

		var Extra_speed_for_flight_to_the_planets = parseFloat(elExtra_speed_for_flight_to_the_planets.value); -- > Orbit


		var Intermediate_angle = Math.abs(Orbit_Inclination) > Math.abs(Spaceport_latitude) ?
				(Rad2Deg(Math.asin(Math.cos(Deg2Rad(Orbit_Inclination))/Math.cos(Deg2Rad(Spaceport_latitude))))) :
				(90.-Rad2Deg(Math.asin(1.-(Math.cos(Deg2Rad(Orbit_Inclination))/Math.cos(Deg2Rad(Spaceport_latitude))))));



		elIntermediate_angle.textContent = Intermediate_angle.toFixed(2);

		var earth_radius = 6371.;
		var gravity_const = 398600500000.;

		var Launch_point_speed = 465.*Math.cos(Deg2Rad(Spaceport_latitude));
		elLaunch_point_speed.textContent = Launch_point_speed.toFixed(1);

		var Starting_point_Altitude_orbital_velocity = Math.sqrt(gravity_const/(earth_radius+Starting_point_Altitude));
		elStarting_point_Altitude_orbital_velocity.textContent = Starting_point_Altitude_orbital_velocity.toFixed(1);

		var Absolute_orbital_velocity = Math.sqrt(gravity_const/(earth_radius+Orbit_Perigee));
		elAbsolute_orbital_velocity.textContent = Absolute_orbital_velocity.toFixed(1);

		var Perigee_velocity = (Orbit_Apogee == 0) ?
			1000.*Math.sqrt(398600.5*(2./(6371.+Orbit_Perigee)-1/(6371.+Orbit_Perigee))) :
			1000.*Math.sqrt(398600.5*(2./(6371.+Orbit_Perigee)-1/(6371.+.5*(Orbit_Perigee+Orbit_Apogee))));
		elPerigee_velocity.textContent = Perigee_velocity.toFixed(1);

		var Apogee_velocity = (Orbit_Apogee == 0) ?
			Perigee_velocity :
			1000.*Math.sqrt(398600.5*(2./(6371.+Orbit_Apogee)-1/(6371.+.5*(Orbit_Perigee+Orbit_Apogee))));
		elApogee_velocity.textContent = Apogee_velocity.toFixed(1);

		var Orbital_period = (Orbit_Apogee == 0) ?
			2.*Math.PI/Math.sqrt(398600.5)*Math.pow(6371+Orbit_Perigee, 3./2.)/60. :
			2.*Math.PI/Math.sqrt(398600.5)*Math.pow(6371+.5*(Orbit_Perigee+Orbit_Apogee), 3./2.)/60.;
		elOrbital_period.textContent = Orbital_period.toFixed(2);

		var Vsp_for_inclination_change = (elOrbit.selectedIndex == 5) ?
//			2*1597.*Math.sin(Math.abs(Deg2Rad(Orbit_Inclination)/2.)) :
			-3.215E-3*Math.pow(Orbit_Inclination,3) + .5134*Math.pow(Orbit_Inclination,2) + .571*Orbit_Inclination :
			0.;
		elVsp_for_inclination_change.textContent = Vsp_for_inclination_change.toFixed(1);

		var tmp_velocity = Math.sqrt(Launch_point_speed*Launch_point_speed + Absolute_orbital_velocity*Absolute_orbital_velocity
										-2.*Launch_point_speed*Absolute_orbital_velocity*Math.sin(Deg2Rad(Intermediate_angle)));

		var Orbital_velocity = (465.*Math.cos(Deg2Rad(Orbit_Inclination)) < Absolute_orbital_velocity) ?
								tmp_velocity : -tmp_velocity;

		tmp_velocity = Math.sqrt(Launch_point_speed*Launch_point_speed + Absolute_orbital_velocity*Absolute_orbital_velocity
										-2.*Launch_point_speed*Absolute_orbital_velocity*Math.sin(Deg2Rad(Intermediate_angle)+1e-7));

		var Orbital_velocity_ = (465.*Math.cos(Deg2Rad(Orbit_Inclination)) < Absolute_orbital_velocity) ?
								tmp_velocity : -tmp_velocity;

		elOrbital_velocity.textContent = Orbital_velocity.toFixed(1);

		var Orbital_velocity_increment_due_to_the_Earth_rotation = Absolute_orbital_velocity - Orbital_velocity;
		elOrbital_velocity_increment_due_to_the_Earth_rotation.textContent = Orbital_velocity_increment_due_to_the_Earth_rotation.toFixed(1);

		var Launch_azimuth = Rad2Deg(Math.acos(Absolute_orbital_velocity/Orbital_velocity*Math.cos(Deg2Rad(Intermediate_angle))));
		var Launch_azimuth_ = Rad2Deg(Math.acos(Absolute_orbital_velocity/Orbital_velocity_*Math.cos(Deg2Rad(Intermediate_angle)+1e-7)));
		if(Launch_azimuth > Launch_azimuth_) Launch_azimuth = -Launch_azimuth;
		elLaunch_azimuth.textContent = Launch_azimuth.toFixed(2);

		Absolute_orbital_velocity0 = Absolute_orbital_velocity;
		if (Absolute_orbital_velocity0 < 7788.5)Absolute_orbital_velocity0 = 7788.5;
		var Vsp_for_circular_orbit0 = Starting_point_Altitude_orbital_velocity
				+Starting_point_Altitude_orbital_velocity*(Starting_point_Altitude_orbital_velocity/Math.sqrt(.5*(Starting_point_Altitude_orbital_velocity*Starting_point_Altitude_orbital_velocity+Absolute_orbital_velocity0*Absolute_orbital_velocity0))-1.)
				+Absolute_orbital_velocity0*(1.-Absolute_orbital_velocity0/Math.sqrt(.5*(Starting_point_Altitude_orbital_velocity*Starting_point_Altitude_orbital_velocity+Absolute_orbital_velocity0*Absolute_orbital_velocity0)));
		Irremovable_Gravity_Losses = - 242.5 + Vsp_for_circular_orbit0 - Absolute_orbital_velocity0;

		var Vsp_for_circular_orbit = Starting_point_Altitude_orbital_velocity
				+Starting_point_Altitude_orbital_velocity*(Starting_point_Altitude_orbital_velocity/Math.sqrt(.5*(Starting_point_Altitude_orbital_velocity*Starting_point_Altitude_orbital_velocity+Absolute_orbital_velocity*Absolute_orbital_velocity))-1.)
				+Absolute_orbital_velocity*(1.-Absolute_orbital_velocity/Math.sqrt(.5*(Starting_point_Altitude_orbital_velocity*Starting_point_Altitude_orbital_velocity+Absolute_orbital_velocity*Absolute_orbital_velocity)))
				-Orbital_velocity_increment_due_to_the_Earth_rotation -Additional_Speed +Vsp_for_inclination_change;
		elVsp_for_circular_orbit.textContent = Vsp_for_circular_orbit.toFixed(1);

		Vsp_for_target_orbit = (Orbit_Apogee == 0) ?
			Vsp_for_circular_orbit +Extra_speed_for_flight_to_the_planets :
			(Vsp_for_circular_orbit +Extra_speed_for_flight_to_the_planets
			  +Absolute_orbital_velocity*(Absolute_orbital_velocity/Math.sqrt(.5*(Absolute_orbital_velocity*Absolute_orbital_velocity+gravity_const/(earth_radius+Orbit_Apogee)))-1.));
		elVsp_for_target_orbit.textContent = Vsp_for_target_orbit.toFixed(1);

		var Injection_velocity = Vsp_for_target_orbit+Orbital_velocity_increment_due_to_the_Earth_rotation;
		elInjection_velocity.textContent = Injection_velocity.toFixed(1);

		var Vsp_for_LEO = 8031.-Orbital_velocity_increment_due_to_the_Earth_rotation-Additional_Speed;
		elVsp_for_LEO.textContent = Vsp_for_LEO.toFixed(1);
	}
*/
