#!/bin/sh -l

command=''
inputFile=''
outputFile=''
output=''
index=''

while getopts c:i:o:t:z: flag
	do
	    case "${flag}" in
	        c) command=${OPTARG};;
	        i) inputFile=${OPTARG};;
	        o) outputFile=${OPTARG};;
	        t) output=${OPTARG};;
	        z) index=${OPTARG};;
	    esac
	done

echo "Command: $command"

echo "inputFile: $inputFile"

if [ -n "$outputFile" ]
  then
    outputFile="--file="$outputFile
    echo "outputFile: $outputFile"
  else
    echo "no outputFile"
fi

if [ -n "$output" ]
  then
    output="--output="$output
    echo "output: $output"
  else
    echo "no output"
fi

if [ -n "$index" ]
  then
    index="--index="$index
    echo "index: $index"
  else
    echo "no index"
fi

/csvbeer "$command" "$inputFile" "$output" "$outputFile" "$index"