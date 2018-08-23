# pigo-openfaas
OpenFaaS function for Pigo face detector.

### Usage
To run the function locally you have to make sure OpenFaaS is up and running. Follow the official documentation for help. https://docs.openfaas.com/

Clone the repository:
```bash
$ git clone https://github.com/esimov/pigo-openfaas
```

#### Build
```bash 
$ faas-cli build -f stack.yml --gateway=http://<GATEWAY-IP>
```

#### Deploy
```bash 
$ faas-cli deploy -f stack.yml --gateway=http://<GATEWAY-IP>
```

You can access the UI on the url provided to `--gateway`. 

![openfaas](https://user-images.githubusercontent.com/883386/44348404-fcef5280-a4a2-11e8-9b9c-1c34acc23d83.png)

### Result
Once the deployment was successfull `pigo-face-detector` will show up in the function list. You have to provide an image URL then hit invoke. The result will be an image with the detected faces marked with a red rectangle.

<p align="center">
<img src="https://user-images.githubusercontent.com/883386/44348484-388a1c80-a4a3-11e8-8d1d-bc529be52ff3.jpg" title="OpenFaaS Pigo Face Detection"/>
</p>

Source image: http://www.balkaninsight.com/uploads/1/images/2016-06-25/e2066a86e25a978d5ab9f7a00b8db076.jpg

The detection time is always around **~0.4s** for the sample image.

## License

This project is under the MIT License. See the LICENSE file for the full license text.
