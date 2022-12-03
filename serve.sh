#!/bin/sh
./wx433 serve wx_433.conf &1>&2 &
rtl_433
pkill wx433