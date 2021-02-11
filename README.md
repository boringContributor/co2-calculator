## Summary
CLI tool written in golang which determines produced co2 values for a car and its covered distance

## Prerequisite
Tested with go1.15.3 darwin/amd64

## Build
First build the tool with
``` bash
go build
```
## Usage
| Argument                | Info                                    | Default | Required |
|-------------------------|-----------------------------------------|---------|----------|
| --help                  | Get information on how to use cli       |         | no       |
| --output                | Specify out of co2 value                | g       | no       |
| --transportation-method | Type of the vehicle                     | -       | yes      |
| --distance              | Covered distance in m or km             | -       | yes      |
| --distance-unit         | Unit of covered distance either m or km | km      | no       |

It is possible to use a command in the following ways (order is not important)
``` bash
--output kg -output kg --output=kg -output=kg 
```
#### Examples
``` bash
./co2-calculator --transportation-method train --distance 14500 --unit-of-distance m --output kg
```

``` bash
./co2-calculator --transportation-method bus --distance 1450 
```

#### Allowed car input

| Small                   | Medium                   | Large                   | Misc. | 
|-------------------------|--------------------------|-------------------------|-----|
| small-diesel-car        | medium-diesel-car        | large-diesel-car        |bus|
| small-petrol-car        | medium-petrol-car        | large-petrol-car        |train|
| small-plugin-hybrid-car | medium-plugin-hybrid-car | large-plugin-hybrid-car ||
| small-electric-car      | medium-electric-car      | large-electric-car      ||

#### Test

Coverage is "only" 40% because there are no tests for the actual parsing of the arguments. This is handled by the flag package itself and only input validations are useful.

``` bash
go test (-cover to see coverage)
```