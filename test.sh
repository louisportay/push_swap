ARGS=`ruby -e "puts (-1000..1000).to_a.shuffle.join(' ')"`

./solver $ARGS | checker $ARGS | (read ret; if [ $ret = "KO" ]; then echo $ARGS > input.txt ; fi )
