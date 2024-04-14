# Obsidian to mela converter

Read our own obsidian template and convert it to a mela file format

## Usage

1. Export recipes to a [`melarecipes` file](https://mela.recipes/fileformat)
```
$> go run ~/path/to/obisidian/recipes ~/path/to/export
recipes exported to "~/path/to/export/mela_1713083022"
$> cd ~/path/to/export
$> zip -r my-recipes.melarecipes mela_1713083022
```
2. Import recipes to some iCloud folder
3. Open "Mela" and "Import recipes" using the `melarecipes` file
