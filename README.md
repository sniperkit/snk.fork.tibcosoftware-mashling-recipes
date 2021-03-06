# Recipes for Project Mashling<sup>TM</sup> 

[![Build Status](https://travis-ci.org/TIBCOSoftware/mashling-recipes.svg?branch=master)](https://travis-ci.org/TIBCOSoftware/mashling-recipes)

A recipe is a pre-configured Mashling json file which can be customized or used as is for a specific gateway use case. The recipes and the gateway binaries to run them are available in [mashling.io](https://www.mashling.io) website.

## How to contribute a recipe

### Creating a recipe
A recipe can be created by customizing an existing recipe in [mashling.io](https://www.mashling.io) or by using an editor such as [VS Code](https://code.visualstudio.com/) enabled with [mashling extension](https://github.com/TIBCOSoftware/vscode-extension-mashling). 

### Adding a recipe
A recipe should be contained in its own folder under 'recipes' folder. The recipe folder is composed of a gateway json file, README.md and optional icon image file. In the absence of the icon image file, the default Mashling icon image is used by mashling.io for the recipe. When the icon image file is present, the Mashling json file should have an icon image file field as follows:

```json
{
	"mashling_schema": "0.2",
	"gateway": {
		"name": "allRecipe",
		"version": "1.0.0",
		"display_name":"KafkaTrigger to KafkaPublisher",
		"display_image":"displayImage.svg",
		"..."
  }
}
```

If "display_name" field is present in the json, its value is used as the recipe name in mashling.io. Otherwise, the value of "name" field is used.

### Publishing a recipe

[recipe_registry.json](https://github.com/TIBCOSoftware/mashling-recipes/blob/master/recipe_registry.json) contains the list of recipe providers and the recipes to publish. The recipe folder name should be added to the "publish" field for the recipe to be made available in mashling.io. For example, "event-dispatcher-router-mashling" and "rest-conditional-gateway" recipes binaries are built and made downloadable from mashling.io given the following recipe_registry.json. Setting "featured" to "true" adds the recipe to the featured recipe list in [mashling.io](https://www.mashling.io).

```json
{  
  "recipe_repos":[  
    {  
      "provider":"TIBCOSoftware Engineering",
      "description":"Mashling gateway recipes from TIBCOSoftware Engineering",
      "publish":[  
        {  
          "recipe":"event-dispatcher-router-mashling",
          "featured":true
        },
        {  
          "recipe":"rest-conditional-gateway",
          "featured":false
        }
      ]
    },
    {  
      "provider":"TIBCOSoftware Services",
      "description":"Mashling gateway recipes from TIBCO Services",
      "publish":[]
    }
  ]
}
```
### Submitting a new/updated recipe
[Create a pull request](https://help.github.com/articles/creating-a-pull-request/) for the recipe to be reviewed and merged into this repository. To publish/unpublish a recipe on [mashling.io](https://www.mashling.io), create a pull request for the updated [recipe_registry.json](https://github.com/TIBCOSoftware/mashling-recipes/blob/master/recipe_registry.json).

## License
mashling-recipes is licensed under a BSD-type license. See license text [here](https://github.com/TIBCOSoftware/mashling-recipes/blob/master/TIBCO%20LICENSE.txt).

## Support
You can post your questions via [GitHub issues](https://github.com/TIBCOSoftware/mashling-recipes/issues).
