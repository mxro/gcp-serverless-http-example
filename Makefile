MAKEFLAGS += -j2

.PHONY: deploy
deploy: deploy-hello deploy-hellov2

.PHONY: deploy-hello
deploy-hello:
	gcloud functions deploy HelloHTTP --runtime=go119 --trigger-http --source=functions/hello --allow-unauthenticated 

.PHONY: deploy-hellov2
deploy-hellov2:
	gcloud functions deploy hello-http-v2 --gen2 --runtime=go119 --source=functionsv2/hello --entry-point=hello-http-v2 --region=us-central1 --trigger-http  --allow-unauthenticated 
