# Team Manager

RESTful API that will help you manage your team. You can create a member (employee or contractor) and attach a tag to him.

### The Task

In this task, we are building the backend of an application that helps us managing our team.

#### Features and Requirements
- A member has a name and a type the late one can be an employee or a contractor - if it's a contractor, the duration of the contract needs to be saved, and if it's an employee we need to store their role, for instance: Software Engineer, Project Manager and so on.
- A member can be tagged, for instance: C#, Angular, General Frontend, Seasoned Leader and so on. (Tags will likely be used as filters later, so keep that in mind)

We need to offer a RESTful CRUD for all the information above.

## Development

If you use VSCode you get started easily using the extension [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers).

Or you can start the project manually running the `docker-compose.yml`:

```
docker-compose -f .devcontainer/docker-compose.yml up -d
```

You can use the application on [http://127.0.0.1:3000](http://127.0.0.1:3000).

### Swagger

When changes are made, you have to init the docs by running the following command on the terminal:

```
swag init --parseDependency --parseInternal -g actions/app.go
```

## How to use the application

Run the application, then open the documentation on http://localhost:3000/v1/doc/index.html. All endpoint are available for test.

## How to deploy

Follow the steps to [install the Convox CLI](https://docsv2.convox.com/introduction/installation).

Once you have a rack up and running, create a new app and deploy the application.

```
$ convox apps create --wait
$ convox deploy --wait
```

To access the aplication, run services and open the URL:

```
$ convox services
SERVICE  DOMAIN                                     PORTS
web      team-manager-web...convox.site  80:3000 443:3000
```

hi test 2
