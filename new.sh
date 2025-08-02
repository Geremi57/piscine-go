curl https://api.github.com/user | jq '.[] | select(.id == 35)'

# echo $request