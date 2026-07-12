#!/bin/bash

setup(){
	export ENV_SOCK_PATH=./test.listen.socket
	
	test -e "${ENV_SOCK_PATH}" && exec env sock="${ENV_SOCK_PATH}" sh -c '
		echo sock "${sock}" already exists.
		echo please remove it to run the example.
		exit 1
	'
	
	./cmd/ustream2discard/ustream2discard &
	
	test -f "${ENV_SOCK_PATH}" || sleep 1
}

bench4linux(){
	setup

	echo longer zeros bench for linux

	time dd if=/dev/zero bs=1048576 count=16384 status=progress |
		nc -N -U "${ENV_SOCK_PATH}"
}

bench4mac(){
	setup

	echo shorter zeros bench for mac

	time dd if=/dev/zero bs=1048576 count=2048 status=progress |
		nc -U "${ENV_SOCK_PATH}"
}

bench(){
	local ostyp
	ostyp=$1
	readonly ostyp

	case "${ostyp}" in
		Linux*)
			bench4linux
			;;

		Darwin*)
			bench4mac
			;;

		*)
			echo "unsupported os type: ${ostyp}"
			exit 1
			;;
	esac
}

bench "$( uname -s )"
