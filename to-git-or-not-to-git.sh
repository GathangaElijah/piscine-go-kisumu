curl -s https://learn.zone01kisumu.ke/assets/superhero/all.json | jq ' map (select(.id == 170))[0] |.name, .powerstats.power, .appearance.gender' 
