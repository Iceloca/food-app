package com.example.foodapp.Management;

import android.content.Context;
import android.util.Log;
import android.widget.Toast;

import com.example.foodapp.Activity.DetailActivity;
import com.example.foodapp.Api.Cart.CartAPI;
import com.example.foodapp.Domain.Cart.CartAddItemRequest;
import com.example.foodapp.Domain.Cart.CartAddItemResponse;
import com.example.foodapp.Domain.Cart.CartDeleteItemRequest;
import com.example.foodapp.Domain.Cart.CartDeleteItemResponse;
import com.example.foodapp.Domain.Cart.Item;
import com.example.foodapp.Domain.Product.ProductResponse;

import java.util.concurrent.CompletableFuture;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class CartManagement {
    private CartAPI cartApi;

    public CartManagement(CartAPI cartAPI) {
        this.cartApi = cartAPI;
    }

    public CompletableFuture<Boolean> addItem(ProductResponse product, String cartId) {
        CompletableFuture<Boolean> future = new CompletableFuture<>();
        Call<CartAddItemResponse> call = cartApi.AddItem(new CartAddItemRequest(new Item(product, 1), cartId));
        call.enqueue(new Callback<CartAddItemResponse>() {
            @Override
            public void onResponse(Call<CartAddItemResponse> call, Response<CartAddItemResponse> response) {
                if (response.isSuccessful() && response.body() != null) {
                    future.complete(true);
                } else {
                    Log.d("error", response.message());
                    future.complete(false);
                }
            }

            @Override
            public void onFailure(Call<CartAddItemResponse> call, Throwable t) {
                Log.d("error", t.toString());
                future.complete(false);
            }
        });
        return future;
    }

    public CompletableFuture<Boolean> deleteItem(String cartId, int itemId) {
        CompletableFuture<Boolean> future = new CompletableFuture<>();
        CartDeleteItemRequest req = new CartDeleteItemRequest(cartId, itemId);
        Call<CartDeleteItemResponse> call = cartApi.DeleteItem(req);
        call.enqueue(new Callback<CartDeleteItemResponse>() {
            @Override
            public void onResponse(Call<CartDeleteItemResponse> call, Response<CartDeleteItemResponse> response) {
                if (response.isSuccessful() && response.body() != null) {
                    future.complete(true);
                } else {
                    Log.d("error", response.message());
                    future.complete(false);
                }
            }

            @Override
            public void onFailure(Call<CartDeleteItemResponse> call, Throwable t) {
                Log.d("error", t.toString());
                future.complete(false);
            }
        });
        return future;
    }
}
