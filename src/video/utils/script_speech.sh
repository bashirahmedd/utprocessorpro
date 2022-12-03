#!/bin/bash
fn_say() 
{ 
    echo  $*
    nc -zw1 google.com 443
    if [ $? -eq 0 ];then
        local IFS=+;/usr/bin/mplayer -ao alsa -really-quiet -noconsolecontrols "http://translate.google.com/translate_tts?ie=UTF-8&client=tw-ob&q=$*&tl=en"; 
    fi
}

# Test Calls for the above functions
#fn_say $*
#fn_say "Hello Hello, how are you?"
#./include_speech.sh "EB is naughty boy. EB is bad boy, he is touching computer. He is disturbing others. Computer is angry. Computer will slap EB on the cheeks. "