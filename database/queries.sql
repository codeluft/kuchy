-- name: GetRecipe :one
SELECT  recipe.*,
        product.*,
        measurement_unit.*
FROM    recipe
JOIN    product ON recipe.product_id = product.id
JOIN    measurement_unit ON product.measurement_unit_id = measurement_unit.id
WHERE   recipe.id = $1
LIMIT   1;
