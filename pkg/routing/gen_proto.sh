#!/bin/bash
dir=`pwd`
rm -fr ${dir}/message
mkdir  ${dir}/message
protoc -I${dir} --go_out=${dir}/message ${dir}/message.proto