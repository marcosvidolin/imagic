# imagic

get the image bytes 

xxd -l 16 -p myimage.jpeg | awk '{gsub(/../,"0x& ",$0); print "["substr($0, 1, length($0)-1)"]"}'
