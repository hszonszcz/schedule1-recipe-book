package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Effect string

const (
	AntiGravity      Effect = "Anti-Gravity"
	Athletic         Effect = "Athletic"
	Balding          Effect = "Balding"
	BrightEyed       Effect = "Bright-Eyed"
	Calming          Effect = "Calming"
	CalorieDense     Effect = "Calorie-Dense"
	Cyclopean        Effect = "Cyclopean"
	Disorienting     Effect = "Disorienting"
	Electrifying     Effect = "Electrifying"
	Energizing       Effect = "Energizing"
	Euphoric         Effect = "Euphoric"
	Explosive        Effect = "Explosive"
	Focused          Effect = "Focused"
	Foggy            Effect = "Foggy"
	Gingeritis       Effect = "Gingeritis"
	Glowing          Effect = "Glowing"
	Jennerising      Effect = "Jennerising"
	Laxative         Effect = "Laxative"
	LongFaced        Effect = "Long Faced"
	Munchies         Effect = "Munchies"
	Paranoia         Effect = "Paranoia"
	Refreshing       Effect = "Refreshing"
	Schizophrenia    Effect = "Schizophrenia"
	Sedating         Effect = "Sedating"
	SeizureInducing  Effect = "Seizure-Inducing"
	Shrinking        Effect = "Shrinking"
	Slippery         Effect = "Slippery"
	Smelly           Effect = "Smelly"
	Sneaky           Effect = "Sneaky"
	Spicy            Effect = "Spicy"
	ThoughtProvoking Effect = "Thought-Provoking"
	Toxic            Effect = "Toxic"
	TropicThunder    Effect = "Tropic Thunder"
	Zombifying       Effect = "Zombifying"
)

type EffectColor struct {
	Name  Effect
	Color string // ANSI color code
}

var effectColors = map[Effect]string{
	AntiGravity:      "\033[38;2;35;91;203m",   // rgb(35, 91, 203)
	Athletic:         "\033[38;2;117;200;253m", // rgb(117, 200, 253)
	Balding:          "\033[38;2;199;146;50m",  // rgb(199, 146, 50)
	BrightEyed:       "\033[38;2;190;247;253m", // rgb(190, 247, 253)
	Calming:          "\033[38;2;254;208;155m", // rgb(254, 208, 155)
	CalorieDense:     "\033[38;2;254;132;244m", // rgb(254, 132, 244)
	Cyclopean:        "\033[38;2;254;193;116m", // rgb(254, 193, 116)
	Disorienting:     "\033[38;2;254;117;81m",  // rgb(254, 117, 81)
	Electrifying:     "\033[38;2;85;200;253m",  // rgb(85, 200, 253)
	Energizing:       "\033[38;2;154;254;109m", // rgb(154, 254, 109)
	Euphoric:         "\033[38;2;254;234;116m", // rgb(254, 234, 116)
	Explosive:        "\033[38;2;254;75;64m",   // rgb(254, 75, 64)
	Focused:          "\033[38;2;117;241;253m", // rgb(117, 241, 253)
	Foggy:            "\033[38;2;176;176;175m", // rgb(176, 176, 175)
	Gingeritis:       "\033[38;2;254;136;41m",  // rgb(254, 136, 41)
	Glowing:          "\033[38;2;133;228;89m",  // rgb(133, 228, 89)
	Jennerising:      "\033[38;2;254;141;248m", // rgb(254, 141, 248)
	Laxative:         "\033[38;2;118;60;37m",   // rgb(118, 60, 37)
	LongFaced:        "\033[38;2;254;217;97m",  // rgb(254, 217, 97)
	Munchies:         "\033[38;2;201;110;87m",  // rgb(201, 110, 87)
	Paranoia:         "\033[38;2;196;103;98m",  // rgb(196, 103, 98)
	Refreshing:       "\033[38;2;178;254;152m", // rgb(178, 254, 152)
	Schizophrenia:    "\033[38;2;100;90;253m",  // rgb(100, 90, 253)
	Sedating:         "\033[38;2;107;95;216m",  // rgb(107, 95, 216)
	SeizureInducing:  "\033[38;2;254;233;0m",   // rgb(254, 233, 0)
	Shrinking:        "\033[38;2;182;254;218m", // rgb(182, 254, 218)
	Slippery:         "\033[38;2;162;223;253m", // rgb(162, 223, 253)
	Smelly:           "\033[38;2;125;188;49m",  // rgb(125, 188, 49)
	Sneaky:           "\033[38;2;123;123;123m", // rgb(123, 123, 123)
	Spicy:            "\033[38;2;254;107;76m",  // rgb(254, 107, 76)
	ThoughtProvoking: "\033[38;2;254;160;203m", // rgb(254, 160, 203)
	Toxic:            "\033[38;2;95;154;49m",   // rgb(95, 154, 49)
	TropicThunder:    "\033[38;2;254;159;71m",  // rgb(254, 159, 71)
	Zombifying:       "\033[38;2;113;171;93m",  // rgb(113, 171, 93)
}

const resetColor = "\033[0m"

type Reagent struct {
	Name    string
	Effects []Effect
	Rules   []ReagentRule
}

type ReagentRule struct {
	Condition   Effect
	Replacement Effect
	RemoveOnly  bool
}

type BaseDrug struct {
	Name    string
	Effects []Effect
}

type DrugMixture struct {
	Base     *BaseDrug
	Reagents []*Reagent
	Effects  map[Effect]bool
	Cost     int
}

type Recipe struct {
	Base     *BaseDrug
	Reagents []*Reagent
	Effects  map[Effect]bool
}

type RecipeStep struct {
	Reagent *Reagent
	Effects []Effect
}

func initializeGameData() ([]*BaseDrug, []*Reagent) {
	bases := []*BaseDrug{
		{Name: "OG Kush", Effects: []Effect{Calming}},
		{Name: "Sour Diesel", Effects: []Effect{Refreshing}},
		{Name: "Green Crack", Effects: []Effect{Energizing}},
		{Name: "Granddady Purple", Effects: []Effect{Sedating}},
		{Name: "Meth", Effects: []Effect{}},
	}

	reagents := []*Reagent{
		{
			Name:    "🥤 Cuke",
			Effects: []Effect{Energizing},
			Rules: []ReagentRule{
				{Condition: Euphoric, Replacement: Laxative},
				{Condition: Foggy, Replacement: Cyclopean},
				{Condition: Gingeritis, Replacement: ThoughtProvoking},
				{Condition: Munchies, Replacement: Athletic},
				{Condition: Slippery, Replacement: Munchies},
				{Condition: Sneaky, Replacement: Paranoia},
				{Condition: Toxic, Replacement: Euphoric},
			},
		},
		{
			Name:    "🧴 Flu Medicine",
			Effects: []Effect{Sedating},
			Rules: []ReagentRule{
				{Condition: Athletic, Replacement: Munchies},
				{Condition: Calming, Replacement: BrightEyed},
				{Condition: Cyclopean, Replacement: Foggy},
				{Condition: Electrifying, Replacement: Refreshing},
				{Condition: Euphoric, Replacement: Toxic},
				{Condition: Focused, Replacement: Calming},
				{Condition: Laxative, Replacement: Euphoric},
				{Condition: Munchies, Replacement: Slippery},
				{Condition: Shrinking, Replacement: Paranoia},
				{Condition: ThoughtProvoking, Replacement: Gingeritis},
			},
		},
		{
			Name:    "⛽ Gasoline",
			Effects: []Effect{Toxic},
			Rules: []ReagentRule{
				{Condition: Disorienting, Replacement: Glowing},
				{Condition: Electrifying, Replacement: Disorienting},
				{Condition: Energizing, Replacement: Euphoric},
				{Condition: Euphoric, Replacement: Spicy},
				{Condition: Gingeritis, Replacement: Smelly},
				{Condition: Jennerising, Replacement: Sneaky},
				{Condition: Laxative, Replacement: Foggy},
				{Condition: Munchies, Replacement: Sedating},
				{Condition: Paranoia, Replacement: Calming},
				{Condition: Shrinking, Replacement: Focused},
				{Condition: Sneaky, Replacement: TropicThunder},
			},
		},
		{
			Name:    "🍩 Donut",
			Effects: []Effect{CalorieDense},
			Rules: []ReagentRule{
				{Condition: AntiGravity, Replacement: Slippery},
				{Condition: Balding, Replacement: Sneaky},
				{Condition: CalorieDense, Replacement: Explosive},
				{Condition: Focused, Replacement: Euphoric},
				{Condition: Jennerising, Replacement: Gingeritis},
				{Condition: Munchies, Replacement: Calming},
				{Condition: Shrinking, Replacement: Energizing},
			},
		},
		{
			Name:    "⚡ Energy Drink",
			Effects: []Effect{Athletic},
			Rules: []ReagentRule{
				{Condition: Disorienting, Replacement: Electrifying},
				{Condition: Euphoric, Replacement: Energizing},
				{Condition: Focused, Replacement: Shrinking},
				{Condition: Foggy, Replacement: Laxative},
				{Condition: Glowing, Replacement: Disorienting},
				{Condition: Schizophrenia, Replacement: Balding},
				{Condition: Sedating, Replacement: Munchies},
				{Condition: Spicy, Replacement: Euphoric},
				{Condition: TropicThunder, Replacement: Sneaky},
			},
		},
		{
			Name:    "💧 Mouth Wash",
			Effects: []Effect{Balding},
			Rules: []ReagentRule{
				{Condition: Calming, Replacement: AntiGravity},
				{Condition: CalorieDense, Replacement: Sneaky},
				{Condition: Explosive, Replacement: Sedating},
				{Condition: Focused, Replacement: Jennerising},
			},
		},
		{
			Name:    "🛢️ Motor Oil",
			Effects: []Effect{Slippery},
			Rules: []ReagentRule{
				{Condition: Energizing, Replacement: Munchies},
				{Condition: Euphoric, Replacement: Sedating},
				{Condition: Foggy, Replacement: Toxic},
				{Condition: Munchies, Replacement: Schizophrenia},
				{Condition: Paranoia, Replacement: AntiGravity},
			},
		},
		{
			Name:    "🍌 Banana",
			Effects: []Effect{Gingeritis},
			Rules: []ReagentRule{
				{Condition: Calming, Replacement: Sneaky},
				{Condition: Cyclopean, Replacement: Energizing},
				{Condition: Disorienting, Replacement: Focused},
				{Condition: Energizing, Replacement: ThoughtProvoking},
				{Condition: Focused, Replacement: SeizureInducing},
				{Condition: LongFaced, Replacement: Refreshing},
				{Condition: Paranoia, Replacement: Jennerising},
				{Condition: Smelly, Replacement: AntiGravity},
				{Condition: Toxic, Replacement: Smelly},
			},
		},
		{
			Name:    "🌶️ Chili",
			Effects: []Effect{Spicy},
			Rules: []ReagentRule{
				{Condition: AntiGravity, Replacement: TropicThunder},
				{Condition: Athletic, Replacement: Euphoric},
				{Condition: Laxative, Replacement: LongFaced},
				{Condition: Munchies, Replacement: Toxic},
				{Condition: Shrinking, Replacement: Refreshing},
				{Condition: Sneaky, Replacement: BrightEyed},
				//{Condition: ThoughtProvoking, Replacement: Focused},
			},
		},
		{
			Name:    "🧪 Iodine",
			Effects: []Effect{Jennerising},
			Rules: []ReagentRule{
				{Condition: Calming, Replacement: Balding},
				{Condition: CalorieDense, Replacement: Gingeritis},
				{Condition: Euphoric, Replacement: SeizureInducing},
				{Condition: Foggy, Replacement: Paranoia},
				{Condition: Refreshing, Replacement: ThoughtProvoking},
				{Condition: Toxic, Replacement: Sneaky},
			},
		},
		{
			Name:    "⚪ Paracetamol",
			Effects: []Effect{Sneaky},
			Rules: []ReagentRule{
				{Condition: Calming, Replacement: Slippery},
				{Condition: Electrifying, Replacement: Athletic},
				{Condition: Energizing, Replacement: Paranoia},
				{Condition: Focused, Replacement: Gingeritis},
				{Condition: Foggy, Replacement: Calming},
				{Condition: Glowing, Replacement: Toxic},
				{Condition: Munchies, Replacement: AntiGravity},
				{Condition: Paranoia, Replacement: Balding},
				{Condition: Spicy, Replacement: BrightEyed},
				{Condition: Toxic, Replacement: TropicThunder},
			},
		},
		{
			Name:    "🔵 Viagra",
			Effects: []Effect{TropicThunder},
			Rules: []ReagentRule{
				{Condition: Athletic, Replacement: Sneaky},
				{Condition: Disorienting, Replacement: Toxic},
				{Condition: Euphoric, Replacement: BrightEyed},
				{Condition: Laxative, Replacement: Calming},
				{Condition: Shrinking, Replacement: Gingeritis},
			},
		},
		{
			Name:    "🐎 Horse Semen",
			Effects: []Effect{LongFaced},
			Rules: []ReagentRule{
				{Condition: AntiGravity, Replacement: Calming},
				{Condition: Gingeritis, Replacement: Refreshing},
				{Condition: SeizureInducing, Replacement: Energizing},
				{Condition: ThoughtProvoking, Replacement: Electrifying},
			},
		},
		{
			Name:    "🥔 Mega Bean",
			Effects: []Effect{Foggy},
			Rules: []ReagentRule{
				{Condition: Athletic, Replacement: Laxative},
				{Condition: Calming, Replacement: Glowing},
				{Condition: Energizing, Replacement: Cyclopean},
				{Condition: Focused, Replacement: Disorienting},
				{Condition: Jennerising, Replacement: Paranoia},
				{Condition: SeizureInducing, Replacement: Focused},
				{Condition: Shrinking, Replacement: Electrifying},
				{Condition: Slippery, Replacement: Toxic},
				{Condition: Sneaky, Replacement: Calming},
				{Condition: ThoughtProvoking, Replacement: Energizing},
			},
		},
		{
			Name:    "💊 Addy",
			Effects: []Effect{ThoughtProvoking},
			Rules: []ReagentRule{
				{Condition: Explosive, Replacement: Euphoric},
				{Condition: Foggy, Replacement: Energizing},
				{Condition: Glowing, Replacement: Refreshing},
				{Condition: LongFaced, Replacement: Electrifying},
				{Condition: Sedating, Replacement: Gingeritis},
			},
		},
		{
			Name:    "🔋 Battery",
			Effects: []Effect{BrightEyed},
			Rules: []ReagentRule{
				{Condition: Cyclopean, Replacement: Glowing},
				{Condition: Electrifying, Replacement: Euphoric},
				{Condition: Euphoric, Replacement: Zombifying},
				{Condition: Laxative, Replacement: CalorieDense},
				{Condition: Munchies, Replacement: TropicThunder},
				{Condition: Shrinking, Replacement: Munchies},
			},
		},
	}

	return bases, reagents
}
func FindShortestRecipe(desired []Effect, base *BaseDrug, reagents []*Reagent) []RecipeStep {
	queue := [][]RecipeStep{
		{
			{
				Reagent: nil,
				Effects: base.Effects,
			},
		},
	}

	visited := make(map[string]bool)
	visited[effectKey(base.Effects)] = true

	for len(queue) > 0 {
		currentPath := queue[0]
		queue = queue[1:]
		currentEffects := currentPath[len(currentPath)-1].Effects

		if containsAll(currentEffects, desired) {
			return currentPath[1:] // Skip the base step in results
		}

		for _, reagent := range reagents {
			newEffects := applyReagent(currentEffects, reagent)

			key := effectKey(newEffects)
			if !visited[key] {
				visited[key] = true
				newPath := append([]RecipeStep(nil), currentPath...)
				newPath = append(newPath, RecipeStep{
					Reagent: reagent,
					Effects: newEffects,
				})
				queue = append(queue, newPath)
			}
		}
	}

	return nil // No recipe found
}

// func applyReagent(current []Effect, reagent *Reagent) []Effect {
// 	effectMap := make(map[Effect]bool)
// 	for _, e := range current {
// 		effectMap[e] = true
// 	}

// 	for _, rule := range reagent.Rules {
// 		if effectMap[rule.Condition] {
// 			delete(effectMap, rule.Condition)
// 			effectMap[rule.Replacement] = true
// 		}
// 	}

// 	for _, e := range reagent.Effects {
// 		effectMap[e] = true
// 	}

// 	var result []Effect
// 	for e := range effectMap {
// 		result = append(result, e)
// 	}
// 	return result
// }

func applyReagent(current []Effect, reagent *Reagent) []Effect {
	// Create a working copy
	newEffects := make(map[Effect]bool)
	for _, e := range current {
		newEffects[e] = true
	}

	// First pass: Process all replacement rules
	replacements := make(map[Effect]Effect)
	for _, rule := range reagent.Rules {
		if newEffects[rule.Condition] {
			replacements[rule.Condition] = rule.Replacement
			delete(newEffects, rule.Condition)
		}
	}

	// Apply all replacements atomically
	for oldEffect, newEffect := range replacements {
		delete(newEffects, oldEffect)
		newEffects[newEffect] = true
	}

	// Add base effects
	for _, e := range reagent.Effects {
		newEffects[e] = true
	}

	// Convert back to slice
	result := make([]Effect, 0, len(newEffects))
	for e := range newEffects {
		result = append(result, e)
	}
	return result
}

func effectKey(effects []Effect) string {
	copied := make([]string, len(effects))
	for i, e := range effects {
		copied[i] = string(e)
	}
	sort.Strings(copied)
	return strings.Join(copied, "|")
}

func containsAll(have []Effect, want []Effect) bool {
	haveMap := make(map[Effect]bool)
	for _, e := range have {
		haveMap[e] = true
	}
	for _, e := range want {
		if !haveMap[e] {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Usage: s1-recipes <efffect1> ...")
		fmt.Println("Example: s1-recipes Spicy Energizing")
		return
	}

	var desired []Effect
	for _, arg := range os.Args[1:] {
		effect := Effect(strings.Title(strings.ToLower(arg)))
		if _, exists := effectColors[effect]; exists {
			desired = append(desired, effect)
		} else {
			fmt.Printf("Warning: Unknown effect '%s' will be ignored\n", arg)
		}
	}

	if len(desired) == 0 {
		fmt.Println("Error: No valid effects specified")
		return
	}

	bases, reagents := initializeGameData()

	fmt.Printf("\n Recipes for ")
	printColoredEffects(desired)

	for _, base := range bases {
		recipe := FindShortestRecipe(desired, base, reagents)
		if recipe != nil {
			printRecipe(base, recipe)
		} else {
			fmt.Printf("No recipe found starting with %s\n", base.Name)
		}
	}
}

func colorizeEffect(e Effect) string {
	if color, ok := effectColors[e]; ok {
		return color + string(e) + resetColor
	}
	return string(e) // Unknown effect fallback
}

func printColoredEffects(effects []Effect) {
	for i, e := range effects {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(colorizeEffect(e))
	}
	fmt.Println()
}

func printRecipe(base *BaseDrug, steps []RecipeStep) {
	fmt.Printf("===================================\n")
	fmt.Printf("\nBase: %s --- Starting effects: ", base.Name)
	printColoredEffects(base.Effects)

	for i, step := range steps {
		fmt.Printf("%d. %s\n", i+1, step.Reagent.Name)
		fmt.Printf("   Effects now: ")
		printColoredEffects(step.Effects)
	}
}
