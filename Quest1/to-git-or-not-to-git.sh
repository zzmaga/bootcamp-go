#!/bin/bash

# Fetch the superhero data for ID 170 and extract name, power, and gender
curl -s https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/superheroes.json | \
  jq -r '.[] | select(.id == 170) | "\(.name)\n\(.powerstats.power)\n\(.appearance.gender)"'

