# Stewart Platform Force Balance


## Description
With the current developments around the NASA Artemis program, there has been an increased interest in the technologies around rocketry and other STEM subjects. One of the key areas around aerodynamics and specifically for this project is wind tunnel testing. Wind tunnels are commonly used in the automotive and aeronautical industries to verify mathematical model. Simulation and analysis of scaled models is an important step in the development of aircraft, vehicles and other machines since it provides aerodynamic performance data that can be used to inform any modifications or improvements. Such improvements can be aimed toward making the aircraft or vehicles more efficient or safer.

Wind tunnels simulate the behaviour of models in the presence of airflow thus allowing us to obtain the components that better define this interaction i.e. forces and moments. A force balance is consequently used in conjunction with wind tunnels to measure these these forces and moments.  Such force balances are made possible by use of sensors for data acquisition by a computer. The force balance can be 3-component or 6-component; external or internal force balance. Whereas the force balances give very accurate results, some are expensive to build and use.This is especially true in low income countries and it limits technological development in these areas.  Moreover, some objects require complex manoeuvre simulations to imitate the actual
movements in air. There is therefore the need for a solution that can achieve dynamic positioning of objects in the wind tunnel while at the same time obtaining force and moment measurements.

Our project intends to bridge this gap by the use of a low cost six degrees-of-freedom (DOF) Stewart platform as a force balance. This type of force balance will be able to position models in the wind tunnel with six degrees of freedom as well as measure six components of aerodynamic forces (drag, lift and side) and moments (pitching, rolling and yawing). The desired effect is to reduce the barrier of entry for such technology and inspire a generation of engineers who can contribute to the space and automotive industries. 

Using the Stewart platform as a force sensor requires the actuators to be locked with zero degrees of freedom. Instrumentation of the Stewart platform legs by use of strain gauges results in a force balance. Four strain gauges are attached to each leg in a Wheatstone bridge configuration, allowing the leg to act as a force sensor. There is an output every time the leg is under axial loading which is first amplified before being read. HX711s are used for analog to digital conversion of the input. Inverse kinematics is used to give various servo angles for the platform positioning. The project uses IBM cloud IoT platform to connect to our designed PCB used to control the platform and take measurements. The data is delivered to our dashboard running on IBM’s code engine which also allows us to send commands to the PCB as well. 

## What's the problem?
The high cost of wind tunnel force balances limits educational access to study areas around the scientific study in aerodynamics for aeronautical and automotive applications. Lowering the cost of force balances would increase access to studies in these fields and inspire students to participate in STEM, especially in low-income areas/countries.


## Video
[Demo video](https://youtu.be/RzLx-dEQbK0)
## Architecture 

## running
### run dashboard
`cd HMI` \
`go run main.go`

## Walkthrough

## Contributing

## Authors
- [Sammy Kerata Oina](https://www.linkedin.com/in/sammy-oina-b1774110b/)
- [Mogire Earl Spencer](https://www.linkedin.com/in/earl-spencer-b03056204//)


