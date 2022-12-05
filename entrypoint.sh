#!/bin/sh

set -e
set -o pipefail

## generate report first
gotest-to-html generate

if [[ -z "$INPUT_DASHBOARD_REPO" ]]; then
    exit 0;
fi

## check which data branch
DATA_BRANCH="${INPUT_DASHBOARD_REPO_DATA_BRANCH:-main}"

## clone dashboard repo
git clone --depth 1 -b ${DATA_BRANCH} https://ghactions:${INPUT_DASHBOARD_REPO_TOKEN}@github.com/${INPUT_DASHBOARD_REPO}.git dashboard-repo

## get dir name
DATE=$(date '+%Y-%m-%d')
mkdir -p dashboard-repo/data/executions/${DATE}_${INPUT_RUN_ID}

## copy html to dir
cp -p ${INPUT_HTML_OUTPUT_FILE} dashboard-repo/data/executions/${DATE}_${INPUT_RUN_ID}/

## copy json to dir
cp -p ${INPUT_JSON_OUTPUT_FILE} dashboard-repo/data/executions/${DATE}_${INPUT_RUN_ID}/

## copt latest report to index.html in root
cp -p ${INPUT_HTML_OUTPUT_FILE} dashboard-repo/index.html

## regenerate dashboard
#TODO

## commit and push
cd dashboard-repo
git config user.email "rajatjindal83@gmail.com"
git config user.name "Rajat Jindal"
git add . && git commit -m "generated report for ${DATE}_${INPUT_RUN_ID}"
git push origin main