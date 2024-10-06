package com.example.foodapp.Activity;

import android.os.Bundle;
import android.view.View;

import androidx.activity.EdgeToEdge;
import androidx.appcompat.app.AppCompatActivity;
import androidx.core.graphics.Insets;
import androidx.core.view.ViewCompat;
import androidx.core.view.WindowInsetsCompat;
import androidx.recyclerview.widget.RecyclerView;

import com.example.foodapp.Domain.Product.ProductResponse;
import com.example.foodapp.R;
import com.example.foodapp.databinding.ActivityFoodListBinding;
import com.fasterxml.jackson.databind.ser.Serializers;

import java.util.ArrayList;

import retrofit2.Call;

public class FoodListActivity extends BaseActivity {

    ActivityFoodListBinding binding;
    private int categoryId;
    private String categoryName;


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        binding = ActivityFoodListBinding.inflate(getLayoutInflater());
        setContentView(binding.getRoot());


    }



    private void getIntentExtra() {
        categoryId = getIntent().getIntExtra("CategoryId", 0);
        categoryName = getIntent().getStringExtra("CategoryName");

        binding.nameCategoryList.setText(categoryName);
        binding.returnButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                finish();
            }
        });
    }

    private void initListFood() {
        Call<ArrayList<ProductResponse>> list =
    }

}