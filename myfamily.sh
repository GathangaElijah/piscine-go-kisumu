#! /bin/bash
curl -s https://learn.zone01kisumu.ke/assets/superhero/all.json | jq -r 'map(select(.id=='$HERO_ID'))[0] .connections.relatives | sub("\n";"\\n")'