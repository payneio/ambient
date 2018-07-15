# Ambient

## Overview

Ambient lets you take control of all of your network devices.

Until this reaches v1, you can expect nothing to work.

Current version: v0.0.0

Planned support (because it is what I have) includes:

- SmartThings Hub
- Broadlink RM Mini 3
- Google ChromeCast
- Google Home
- Bosch PIR Motion Detector (via Smart Things Hub)
- SmartPower Outlet (via Smart Things Hub)
- SmartSense Multi-Sensor (via Smart Things Hub)
- Zooz Z-Wave Relay (via Smart Things Hub)
- Wake on LAN for local MAC addresses

## Components

- Discovery routines
- Registration
  - See sensors
  - See actuators
  - See cmd tree
- PubSub via ...
- Events feed
- Sensor metrics aggregation
- Sensor visualization
- Command
  - Allows endpoints for each available effector command
  - Handles execution of command, ensuring it is received and monitoring state as appropriate to ensure command completes successfully
- Reflex
  - Allows the registration of IfTTT-style rules
  - Continuously monitors system state to carry out rules as needed

## Next Milestone

- SmartThings and Broadlink device discovery

