#!/bin/bash

cat ./testdata/yesterday.json | ./jedit addKey "incident_id" "6502" --filter "team == team-x"