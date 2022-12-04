#!/bin/bash -l

set -o pipefail

## generate report first
gotest-to-html

if [[ -z "$INPUT_DASHBOARD_REPO" ]]; then
    exit 0;
fi

## clone dashboard repo
git clone -d 1 https://ghactions:${DASHBOARD_REPO_TOKEN}@github.com/${INPUT_DASHBOARD_REPO}.git dashboard-repo

## get dir name
DATE=$(date '+%Y-%m-%d')
mkdir -p ${INPUT_HTML_OUTPUT_FILE} dashboard-repo/data/executions/${DATE}_${INPUT_RUN_ID}

## copy html to dir
cp -p ${INPUT_HTML_OUTPUT_FILE} dashboard-repo/data/executions/${DATE}_${INPUT_RUN_ID}/

## copy json to dir
cp -p ${INPUT_JSON_OUTPUT_FILE} dashboard-repo/data/executions/${DATE}_${INPUT_RUN_ID}/

## regenerate dashboard
#TODO

## commit and push
cd dashboard-repo
git config user.email "rajatjindal83@gmail.com"
git config user.name "Rajat Jindal"
git add . && git commit -m "generated report for ${DATE}_${INPUT_RUN_ID}"
git push origin main