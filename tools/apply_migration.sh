#!/usr/bin/env bash

migrate -database postgresql://localhost:5432/track?sslmode=disable -path ../backend/_migrations up