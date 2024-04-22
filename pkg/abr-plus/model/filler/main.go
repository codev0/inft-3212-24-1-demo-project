package filler

import (
	model "github.com/codev0/inft3212-6/pkg/abr-plus/model"
)

func PopulateDatabase(models model.Models) error {
	for _, menu := range menus {
		models.Menus.Insert(&menu)
	}
	// TODO: Implement restaurants pupulation
	// TODO: Implement the relationship between restaurants and menus
	return nil
}

var menus = []model.Menu{
	{Title: "Caesar Salad", Description: "Classic Caesar with homemade dressing", NutritionValue: 170},
	{Title: "Greek Salad", Description: "Traditional Greek salad with feta cheese", NutritionValue: 200},
	{Title: "Caprese Salad", Description: "Fresh tomatoes and mozzarella slices", NutritionValue: 180},
	{Title: "Cobb Salad", Description: "Loaded Cobb salad with grilled chicken", NutritionValue: 350},
	{Title: "Kale Salad", Description: "Kale with cranberries and pine nuts", NutritionValue: 220},
	{Title: "Quinoa Salad", Description: "Quinoa with mixed vegetables", NutritionValue: 190},
	{Title: "Spinach Salad", Description: "Spinach with walnuts and balsamic", NutritionValue: 210},
	{Title: "Fruit Salad", Description: "Seasonal fresh fruit mix", NutritionValue: 120},
	{Title: "Pasta Salad", Description: "Chilled pasta with Italian dressing", NutritionValue: 250},
	{Title: "Potato Salad", Description: "Creamy potato salad with herbs", NutritionValue: 280},
	{Title: "Taco Salad", Description: "Spicy taco fillings over greens", NutritionValue: 300},
	{Title: "Chicken Salad", Description: "Chicken salad with celery and mayo", NutritionValue: 230},
	{Title: "Seafood Salad", Description: "Mixed seafood with tangy dressing", NutritionValue: 260},
	{Title: "Egg Salad", Description: "Chopped eggs with creamy dressing", NutritionValue: 240},
	{Title: "Asian Salad", Description: "Asian greens with sesame soy dressing", NutritionValue: 150},
	{Title: "Tuna Salad", Description: "Tuna with crunchy celery and onions", NutritionValue: 250},
	{Title: "Broccoli Salad", Description: "Raw broccoli with bacon bits", NutritionValue: 290},
	{Title: "Carrot Salad", Description: "Shredded carrots with raisins", NutritionValue: 130},
	{Title: "Beet Salad", Description: "Roasted beets with feta", NutritionValue: 160},
	{Title: "Arugula Salad", Description: "Arugula with parmesan shavings", NutritionValue: 140},
	{Title: "Southwest Salad", Description: "Beans and corn with a spicy kick", NutritionValue: 320},
	{Title: "Waldorf Salad", Description: "Apples and walnuts with a creamy dressing", NutritionValue: 270},
	{Title: "Cucumber Salad", Description: "Thinly sliced cucumbers with dill", NutritionValue: 110},
	{Title: "Avocado Salad", Description: "Sliced avocado over mixed greens", NutritionValue: 230},
	{Title: "Mediterranean Salad", Description: "Chickpeas and olives with a lemon dressing", NutritionValue: 290},
	{Title: "Chef Salad", Description: "Ham, turkey, cheese, and eggs over greens", NutritionValue: 330},
	{Title: "Lentil Salad", Description: "Lentils with veggies in a vinaigrette", NutritionValue: 180},
	{Title: "Bacon Salad", Description: "Crispy bacon over spinach and tomatoes", NutritionValue: 310},
	{Title: "Buffalo Chicken Salad", Description: "Spicy buffalo chicken pieces with blue cheese", NutritionValue: 360},
	{Title: "Antipasto Salad", Description: "Italian meats with olives and peppers", NutritionValue: 340},
}
