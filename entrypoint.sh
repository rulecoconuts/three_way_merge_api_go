ls -R
cd website
../three_way_merge & busybox httpd -f -v -p 3000 && fg