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
After deploying the OpenFaaS function `pigo-face-detector` will show up in the function list. You have to provide an image URL then hit invoke. This will return an image with yellow rectangles drawn around the detected faces.

<p align="center">
<img src="https://user-images.githubusercontent.com/883386/53553799-202c4600-3b47-11e9-91d5-ff8f296fbeb3.jpg" title="OpenFaaS Pigo Face Detection"/>
</p>

Sample image: https://user-images.githubusercontent.com/883386/53553708-ebb88a00-3b46-11e9-9ea8-73c6b7f9dfa1.jpg

The detection time is always around **~0.4s** for the sample image.

## License

This project is under the MIT License. See the LICENSE file for the full license text.
