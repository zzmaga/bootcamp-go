subject_id=70
response=$(curl -s https://01.tomorrow-school.ai/assets/superhero/all.json)
echo "$response" | tr -d '\000-\031' | jq -r --arg id "$subject_id" '.[] | select(.id == ($id | tonumber)) | .name' | sed 's/^/"/;s/$/"/'

