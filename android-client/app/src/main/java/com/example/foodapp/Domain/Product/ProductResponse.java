package com.example.foodapp.Domain.Product;

import com.fasterxml.jackson.annotation.JsonProperty;

public class ProductResponse {

    @JsonProperty("id")
    Integer id;

    @JsonProperty("name")
    String name;

    @JsonProperty("description")
    String description;

    @JsonProperty("image_url")
    String imageURL;

    @JsonProperty("price")
    Float price;

    @JsonProperty("is_daily_rec")
    boolean isDailyRec;

    @JsonProperty("category_id")
    int categoryId;

    public Integer getId() {
        return id;
    }

    public String getName() {
        return name;
    }

    public String getDescription() {
        return description;
    }

    public String getImageURL() {
        return imageURL;
    }

    public Float getPrice() {
        return price;
    }

    public int getCategoryId() {
        return categoryId;
    }

    public boolean isDailyRec() {
        return isDailyRec;
    }
}
