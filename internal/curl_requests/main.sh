#!/bin/zsh
P=$(pwd)
curl -X POST http://127.0.0.1:5053/placements/request \
     -H "Content-Type: application/json" \
     -d @simple.json

curl -X POST http://127.0.0.1:5053/placements/request \
     -H "Content-Type: application/json" \
     -d @skip.json

curl -X POST http://127.0.0.1:5053/placements/request \
     -H "Content-Type: application/json" \
     -d @no_resp_for_imp.json
