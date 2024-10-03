package com.example.foodapp.Api.Product;

import com.example.foodapp.Domain.Product.CategoryResponse;
import com.example.foodapp.Domain.Product.ProductResponse;

import java.util.ArrayList;

import retrofit2.Call;
import retrofit2.http.GET;

public interface ProductAPI {
    @GET("api/v1/product")
    Call<ArrayList<ProductResponse>> GetAllProducts();
    @GET("api/v1/product/recs")
    Call<ArrayList<ProductResponse>> GetBestProducts();
    @GET("api/v1/category")
    Call<ArrayList<CategoryResponse>> GetCategories();
}
