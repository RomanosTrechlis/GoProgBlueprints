#!/bin/bash
echo Building synonyms...
cd ../synonyms
go build -o ../domainfinder/lib/synonyms.exe
echo Building available...
cd ../available
go build -o ../domainfinder/lib/available.exe
echo Building sprinkle...
cd ../sprinkle
go build -o ../domainfinder/lib/sprinkle.exe
echo Building coolify...
cd ../coolify
go build -o ../domainfinder/lib/coolify.exe
echo Building domainify...
cd ../domainify
go build -o ../domainfinder/lib/domainify.exe
echo Building domainfinder...
cd ../domainfinder
go build -o domainfinder.exe
echo Done.