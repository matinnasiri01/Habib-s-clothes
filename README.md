# 👔 Habib's Outfit Combinations

The costume designer of the series *Lisansa* is going on a short trip and has given **Habib** the opportunity to choose his outfits according to his own taste during this period.  
Habib has several clothes in his wardrobe and wants to create different combinations.  

For example:

- One day he might wear a brown coat, yellow shirt, and blue pants  
- Another day he might wear a red shirt with black pants (without a coat)

This program will generate **all possible valid outfit combinations** for Habib based on his rules.

---

## 📥 Input

There are 6 lines of input, each representing a category of Habib's clothes:

1. **COAT** – list of coat colors  
2. **SHIRT** – list of shirt colors  
3. **PANTS** – list of pants colors  
4. **CAP** – list of sun hat colors  
5. **JACKET** – list of jacket colors  
6. **SEASON** – one of: `SPRING`, `SUMMER`, `FALL`, `WINTER`

**Example:**

```plaintext
COAT: white brown yellow
SHIRT: yellow pink red
PANTS: white carbon-blue
CAP: purple grey
JACKET: indigo orange
SUMMER
```

---

## 📤 Output

Print all valid outfit combinations according to the rules.

- The order of combinations does not matter.  
- The order of items in each combination must be:
```
COAT SHIRT PANTS CAP JACKET
```
- Do not print items that are not worn.

  **Example:**

```plaintext
SHIRT: yellow PANTS: white CAP: purple
SHIRT: yellow PANTS: white CAP: grey
SHIRT: yellow PANTS: carbon-blue CAP: purple
SHIRT: yellow PANTS: carbon-blue CAP: grey
SHIRT: pink PANTS: white CAP: purple
SHIRT: pink PANTS: white CAP: grey
SHIRT: pink PANTS: carbon-blue CAP: purple
SHIRT: pink PANTS: carbon-blue CAP: grey
SHIRT: red PANTS: white CAP: purple
SHIRT: red PANTS: white CAP: grey
SHIRT: red PANTS: carbon-blue CAP: purple
SHIRT: red PANTS: carbon-blue CAP: grey
```

---

## 📏 Rules

1. Shirt and pants are always worn in all seasons.

2. Cap: must be worn in SUMMER; may be worn in SPRING and FALL.

3. Coat: may be worn in SPRING, FALL, and WINTER.

4. Jacket: may be worn only in WINTER.

5. In WINTER, exactly one of coat or jacket must be worn.

6. In FALL, coat cannot be yellow or orange.

---

## 📝 Examples
Input:
```
COAT: white brown yellow
SHIRT: yellow pink red
PANTS: white carbon-blue
CAP: purple grey
JACKET: indigo orange
SUMMER
```
Output:
```
SHIRT: yellow PANTS: white CAP: purple
SHIRT: yellow PANTS: white CAP: grey
SHIRT: yellow PANTS: carbon-blue CAP: purple
SHIRT: yellow PANTS: carbon-blue CAP: grey
SHIRT: pink PANTS: white CAP: purple
SHIRT: pink PANTS: white CAP: grey
SHIRT: pink PANTS: carbon-blue CAP: purple
SHIRT: pink PANTS: carbon-blue CAP: grey
SHIRT: red PANTS: white CAP: purple
SHIRT: red PANTS: white CAP: grey
SHIRT: red PANTS: carbon-blue CAP: purple
SHIRT: red PANTS: carbon-blue CAP: grey
```
