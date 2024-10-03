package com.example.foodapp.Domain.Product;

import com.fasterxml.jackson.annotation.JsonProperty;

public class CategoryResponse {

    @JsonProperty("id")
    Integer id;

    @JsonProperty("name")
    String name;

    @JsonProperty("image_url")
    String imageURL;

    public Integer getId() {
        return id;
    }

    public String getImageURL() {
        return imageURL;
    }

    public String getName() {
        return name;
    }
}
