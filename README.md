API Documentation

This document provides an overview of all the available API routes for managing equipment, materials, and product specifications in your system. Each route is described with its purpose and the corresponding HTTP method.

Equipment Routes

1. Create Equipment

Route: /api/equipment

Method: POST

Description: Allows the creation of a new equipment item in the system.

2. Get Filtered Equipment

Route: /api/equipment/filter

Method: GET

Description: Retrieves a list of equipment items based on applied filters.

3. Get Equipment with Materials

Route: /api/equipment/materials

Method: GET

Description: Fetches equipment items along with their associated materials.

4. Update Equipment Lifetime

Route: /api/equipment/update-lifetime

Method: PUT

Description: Updates the lifetime duration for a specific equipment item.

5. Delete Equipment

Route: /api/equipment/:id

Method: DELETE

Description: Deletes an equipment item by its ID.

Material Routes

1. Create Material

Route: /api/material

Method: POST

Description: Adds a new material to the system.

2. Get Filtered Materials

Route: /api/material

Method: GET

Description: Retrieves materials based on specified filters.

3. Update Material Price

Route: /api/material/update-price

Method: PUT

Description: Updates the price of a specific material.

Product Specification Routes

1. Create Product Specification

Route: /api/product-specification

Method: POST

Description: Creates a new product specification in the system.

2. Get Specifications by Equipment ID

Route: /api/product-specification/equipment/:id

Method: GET

Description: Retrieves product specifications associated with a specific equipment ID.

3. Update Production Duration

Route: /api/product-specification/:id

Method: PUT

Description: Updates the production duration for a specific product specification by its ID.