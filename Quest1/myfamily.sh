#!/bin/bash
CURL='https://01.tomorrow-school.ai/assets/superhero/all.json'
FAM=$(curl -s $CURL | jq -r --arg id "$HERO_ID" '.[] | select(.id ==($id | tonumber)) | .connections.relatives | gsub ("\n"; "\\n")')
printf "%s" "$FAM"