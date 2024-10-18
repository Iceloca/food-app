package com.example.foodapp.Activity;

import android.annotation.SuppressLint;
import android.content.SharedPreferences;
import android.graphics.Color;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.view.Window;
import android.view.WindowManager;
import android.widget.Toast;

import androidx.activity.EdgeToEdge;
import androidx.appcompat.app.AppCompatActivity;
import androidx.core.content.ContextCompat;
import androidx.core.graphics.Insets;
import androidx.core.view.ViewCompat;
import androidx.core.view.WindowInsetsCompat;

import com.bumptech.glide.Glide;
import com.example.foodapp.Domain.Cart.CartAddItemRequest;
import com.example.foodapp.Domain.Cart.CartAddItemResponse;
import com.example.foodapp.Domain.Cart.Item;
import com.example.foodapp.Domain.Product.ProductResponse;
import com.example.foodapp.Management.CartManagement;
import com.example.foodapp.R;
import com.example.foodapp.databinding.ActivityDetailBinding;
import com.fasterxml.jackson.databind.ser.Serializers;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class DetailActivity extends BaseActivity {

    ActivityDetailBinding binding;
    ProductResponse product;
    String cartId;
    CartManagement cartManagement;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        binding = ActivityDetailBinding.inflate(getLayoutInflater());

        Window window = getWindow();
        window.clearFlags(WindowManager.LayoutParams.FLAG_TRANSLUCENT_STATUS);
        window.addFlags(WindowManager.LayoutParams.FLAG_DRAWS_SYSTEM_BAR_BACKGROUNDS);
        window.setStatusBarColor(Color.WHITE); // Белый цвет
        View decorView = window.getDecorView();
        decorView.setSystemUiVisibility(View.SYSTEM_UI_FLAG_LIGHT_STATUS_BAR);

        setContentView(binding.getRoot());

        cartManagement = new CartManagement(cartApi);

        SharedPreferences cartSharedPreferences = getSharedPreferences("cart_prefs", MODE_PRIVATE);
        SharedPreferences authSharedPreferences = getSharedPreferences("auth_prefs", MODE_PRIVATE);
        int userId = authSharedPreferences.getInt("user_id", -1);
        this.cartId = cartSharedPreferences.getString("cart_id_" + userId, null);

        getIntentExtra();
        setVariable();
    }

    @SuppressLint("SetTextI18n")
    private void setVariable() {

        binding.returnButtonDetail.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                finish();
            }
        });

        binding.addToCartButtonDetail.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                cartManagement.addItem(product, cartId).thenAccept(success -> {
                   if (success) {
                       Toast.makeText(DetailActivity.this, "Блюдо успешно добавлено в корзину заказа", Toast.LENGTH_SHORT).show();
                   }
                });
            }
        });


        binding.nameProductDetail.setText(product.getName());
        binding.descriptionProductDetail.setText(product.getDescription());
        binding.priceProductDetail.setText(product.getPrice() + "BYN");
        binding.timeProductDetail.setText("20мин");

        Glide.with(DetailActivity.this)
                .load(product.getImageURL())
                .into(binding.imageProductDetail);
    }

    private void getIntentExtra() {
        product = (ProductResponse) getIntent().getSerializableExtra("product");
    }
}