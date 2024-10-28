package com.example.foodapp.Api.Cart;

import com.example.foodapp.Domain.Cart.CartAddItemRequest;
import com.example.foodapp.Domain.Cart.CartAddItemResponse;
import com.example.foodapp.Domain.Cart.CartCreateRequest;
import com.example.foodapp.Domain.Cart.CartCreateResponse;
import com.example.foodapp.Domain.Cart.CartDeleteItemRequest;
import com.example.foodapp.Domain.Cart.CartDeleteItemResponse;
import com.example.foodapp.Domain.Cart.CartResponse;

import retrofit2.Call;
import retrofit2.http.Body;
import retrofit2.http.GET;
import retrofit2.http.POST;
import retrofit2.http.PUT;
import retrofit2.http.Path;

public interface CartAPI {
    @GET("api/v1/cart/{id}")
    Call<CartResponse> GetById(@Path("id") String id);
    @POST("api/v1/cart/")
    Call<CartCreateResponse> Create(@Body CartCreateRequest req);
    @PUT("api/v1/cart/item/add")
    Call<CartAddItemResponse> AddItem(@Body CartAddItemRequest req);
    @PUT("api/v1/cart/item/delete")
    Call<CartDeleteItemResponse> DeleteItem(@Body CartDeleteItemRequest req);
}
