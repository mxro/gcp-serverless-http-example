


deploy:
	gcloud functions deploy HelloHTTP --runtime go119 --trigger-http --source=functions/hello --allow-unauthenticated 

deployV2:
	gcloud functions deploy hello-http-v2 --gen2 --runtime=go119 --source=functionsv2/hello --entry-point=hello-http-v2 --region=us-central1 --trigger-http  --allow-unauthenticated 
