# Final Year Report
## Outline 
0. ~~ Title page ~~
1. Abstract
2. Introduction (Chapter 1)
3. Literature Review (Chapter 2)
4. Methodology (Chapter 3)
5. Results and Discussion (Chapter 4)
6. Conclusion (Chapter 5)

**Problem Statement**

We need A STATEMENT for our problem STATEMENT

The force balances used in wind tunnel testing have to incorporate linear actuators in order to alter model position. This, however, only changes the model's angle of attack and dynamic orientation of models is impossible to achieve. 

Line 44 to talk about servos actuating the Stewart platform movements

**Smoke Visualization Technique for Wind Tunnel**

The ability to visualise the airflow over solid objects in a wind tunnel crucial in understanding the fundamental principles of both fluid dynamics and aerodynamics. To achieve this a uniform row of smoke lines are required. These smoke lines should be sufficiently long to maintain their integrity throughout the test section.

Smoke visualization methods including smoke-generating materials and techniques for generating smoke lines have been summarized in the tables below.

The application smoke lines over several test objects, such as wing sections and bluff body shapes, ensures high-quality visualisations.

Carbon dioxide
Produces dense smoke
Potentially harmful in large volumes;
Non-hazardous; produce dense vapour
Vapour can condense back to liquid

Smoke wire
Produces the most effective smoke lines
Can be utilised with non-hazardous smoke materials

For smoke rake, an aerodynamically shaped
body (typically elliptical) featuring a row of tubes through which the smoke exits is used. Smoke to the rake can be introduced from a non-hazardous source, such as a water-based liquid heated by a smoke machine, rather than using combustion of hydrocarbons as in the smoke wire technique, which is hazardous and produces toxic materials.

## Methodology

### Instrumentation of the Stewart Platform

From FEA done on the Stewart platform, strain values were expected along the Stewart platform legs as shown in figure \ref{eq}.

## Electromechanical Modifications

To implement streamline smoke lines in the wind tunnel, a smoke rake, which is an aerodynamically shaped
body (typically elliptical) featuring a row of tubes through which the smoke exits was designed. The CAD model was deigned using Autodesk Inventor and the rake was prepared for 3D printing at iPIC in JKUAT. The rake design is shown below.

To implement smoke visualisation in the wind tunnel, the rake is used in conjuction with:

## Smoke rake
Change top hole diameter to 20mm

## CFD Parameters
1. Simulation - Fluid Flow Fluent
2. Simulation type - steady flow
3. Diameter of strut - 12.5mm
4. Area - 
5. Mesh size - 0.001mm
6. Growth rate - default (1.2)
7. Behavior - soft
8. Defeature size - 5.5902 e-004m

## Specific Objectives
To develop a force balance for the Stewart platform and obtain forces and moments
measurement during model testing.

To obtain flow velocity measurement in the wind tunnel by use of pitot tubes.

To develop a Human-Machine Interface for measurement readings and control of
the Stewart platform load balance.

To improve model behaviour visualization in the wind tunnel by introduction of
smoke streamlines in the wind tunnel