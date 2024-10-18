package com.example.foodapp.Adapter;

import android.annotation.SuppressLint;
import android.content.Context;
import android.content.Intent;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.bumptech.glide.Glide;
import com.bumptech.glide.load.resource.bitmap.CenterCrop;
import com.bumptech.glide.load.resource.bitmap.RoundedCorners;
import com.example.foodapp.Activity.DetailActivity;
import com.example.foodapp.Domain.Product.ProductResponse;
import com.example.foodapp.R;

import java.util.ArrayList;

public class BestProductAdapter extends RecyclerView.Adapter<BestProductAdapter.viewholder> {
    private final ArrayList<ProductResponse> items;
    private Context context;

    public BestProductAdapter(ArrayList<ProductResponse> items) {
        this.items = items;
    }

    @NonNull
    @Override
    public BestProductAdapter.viewholder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        context = parent.getContext();
        View inflate = LayoutInflater.from(parent.getContext()).inflate(R.layout.viewholder_bestproduct, parent, false);
        return new viewholder(inflate);
    }

    @SuppressLint("SetTextI18n")
    @Override
    public void onBindViewHolder(@NonNull BestProductAdapter.viewholder holder, @SuppressLint("RecyclerView") int position) {
        holder.nameBestProduct.setText(items.get(position).getName());
        holder.priceBestProduct.setText(items.get(position).getPrice()+"BYN");
        holder.timeBestProduct.setText("20мин");

        Glide.with(context)
                .load(items.get(position).getImageURL())
                .transform(new CenterCrop(), new RoundedCorners(30))
                .into(holder.imageBestProduct);

        holder.itemView.setOnClickListener(view -> {
            Intent intent = new Intent(context, DetailActivity.class);
            intent.putExtra("product", items.get(position));
            context.startActivity(intent);
        });
    }

    @Override
    public int getItemCount() {
        return items.size();
    }

    public static class viewholder extends RecyclerView.ViewHolder {
        TextView nameBestProduct, priceBestProduct, timeBestProduct;
        ImageView imageBestProduct;
        public viewholder(@NonNull View itemView) {
            super(itemView);
            nameBestProduct = itemView.findViewById(R.id.nameBestProduct);
            timeBestProduct = itemView.findViewById(R.id.timeBestProduct);
            priceBestProduct = itemView.findViewById(R.id.priceBestProduct);
            imageBestProduct = itemView.findViewById(R.id.imageBestProduct);
        }
    }
}
