# D0020E

## Introduction

This Orchestrator is part of the arrowhead framework rewritten in GOlang for the project course D0020E at LTU.

The framework is done in collaboration with two other groups responsible for the Service Registry and the Consumer-Provider part of the framework.

Link to Service Registry: https://github.com/ClaudeHallard/Arrowhead

Link to Consumer-Provider system: https://github.com/DavidS1998/D00020E

The original Arrowhead framework written in Java: https://github.com/eclipse-arrowhead/core-java-spring


## Compile

To compile the orchestrator, simply excecute the main file. Remember to set up the address to the Service Registry in the config file.

## Config

The config file, located in the folder "Orchestrator", includes the address and port used to reach the Orchestrator. As well as the address used by the orchestrator to reach the Service Registry.

**SERVICE_ADDRESS** specifies the address to the Service Registry.

**CONN_PORT** is the port used to reach the Orchestrator.