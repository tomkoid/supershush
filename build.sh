#!/usr/bin/env bash

if [ ! -d "build" ]
then
	mkdir build
fi

package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package>"
  exit 1
fi
package_split=(${package//\// })
package_name=${package_split[-1]}
	
platforms=("linux/amd64" "linux/arm" "linux/ppc64" "linux/ppc64le" "linux/mips" "linux/mipsle" "linux/mips64" "linux/mips64le" "linux/arm64" "linux/386" "freebsd/amd64" "freebsd/arm" "freebsd/386" "netbsd/386" "netbsd/amd64" "netbsd/arm" "openbsd/386" "openbsd/amd64" "openbsd/arm")

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name=supershush-$GOOS'-'$GOARCH
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi	

	echo "Building $GOOS/$GOARCH"
	env GOOS=$GOOS GOARCH=$GOARCH go build -o build/$output_name $package
	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
done
