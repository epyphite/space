# Epyphite's NASA GO API integration

This is a collaborative effort to provide access for GO tool chain to the NASA API and services exposed [here](https://api.nasa.gov/)

## Integrations completed

- [X] APOD (Astronomical Picture Of the Day)
- [X] EONET (Earth Observatory Natural Event Tracker)
- [X] NEO (Near Earth Objects)
- [X] TLE API ( two line element )

## Goal

To build a satellite tracking library tool for golang compatible with gRPC and ResilientOne format and allowing common:

- User and machine friendly interface to interact with
- Easy access for satellite data
- Easiy Integration with third party tools
- Complete separation of input and output data
- Satellite orbit calculation.

## Features

1. [X] Reading of the NORAD Two-Line Element Set Format for the chosen satellite

2. [X] extraction of necessary information such as:
    - satellite name,
    - eccentricity,
    - inclination,
    - argument of periapsis,
    - mean anomaly,
    - longitude of the ascending node and mean motion;

3. [ ] mathematical calculations of necessary input for the application;

4. [ ] starting point determination, all points of the satellite track being characterized by vector radius and vector speed;

5. [ ] Calculatation to obtain the satellite’s orbit;

6. [ ] projection of satellite’s orbit onto Earth’s surface, without and with revolution movement;

7. [ ] Satellite’s rate of decay;

8. [ ] API display in Earth’s atmosphere for the last two steps;
