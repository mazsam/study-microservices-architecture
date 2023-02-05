package demo_go_swagger_api;

//go:generate swagger generate server --exclude-main -A gateway -t gen -f ./api/swagger.yml --principal models.Principal
//go:generate swagger -q generate client -A gateway -f api/swagger.yml -c pkg/client -m gen/models --principal models.Principal
