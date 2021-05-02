# pigo-openfaas

[Pigo](https://github.com/esimov/pigo) OpenFaaS function for face detection.

### Usage
To run the function locally you have to make sure OpenFaaS is up and running. Read the official documentation for more information. https://docs.openfaas.com/

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

## License

Copyright © 2018 Endre Simo

This project is under the MIT License. See the LICENSE file for the full license text.
