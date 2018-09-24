# Quantified Self
is a project designed to teach how seperate backend and frontend applications interact.  Given a basic frontend application written in React (https://github.com/Vadlusk/quantified-self-fe-go), students are tasked with creating three matching backends, the first in Rails, the second in Express, and the third in a language of our choice.  
The application's functionality is that of a basic calorie counter.  The user can add foods to a list and then add foods from that list to any of four meals.  Each food is assigned a number of calories which is subtracted from remaining calories for each meal and for the day (all meals added together).
You can visit the [deployed backend here](https://safe-spire-34566.herokuapp.com/api/v1/meals/), or the [deployed frontend here](https://vadlusk.github.io/quantified-self-fe-go/).

# Versions

# Setup
1. Clone or fork this repo in your go workspace.
2. Move into the directory.
3. `createdb quantified_self_go_dev`
4. `go build`
5. `./quantified_self_be_go`

# Endpoints

## Foods
The following endpoints return either a food object or message that the action was successful.  
Food objects are returned in the following format:   
```
{
  "id": int,
  "name": string,
  "calories": int
}
```

GET /api/v1/foods

returns a list of all foods in the database

GET /api/v1/foods/:id

returns the food with that id number

POST /api/v1/foods

creates a food with paramaters:
```
{
  "food": {
    "name": string,
    "calories": int
  }
}
```

all fields are required

PATCH /api/v1/foods/:id

edits an existing food with same parameters as above.   
no fields are required.

DELETE /api/v1/foods/:id

deletes the food with that id number

## Meals
The following endpoints return either a meal object or message that the action was successful.  
Meal objects are returned in the following format:   
```
{
  "id": int,
  "name": string,
  "foods": array_of_food_objects
}
```

GET /api/v1/meals

returns all the meals in the database with associated foods

GET /api/v1/meals/:meal_id/foods

returns the meal with that id number

POST /api/v1/meals/:meal_id/foods/:id

associates the food with that id with the meal with that id

DELETE /api/v1/meals/:meal_id/foods/:id

removes the association between that food and that meal
