package api

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"github.com/k2hmr/panforyou-test-1/internal/model"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func SaveBread(saveData model.SaveData, ctx context.Context) error {
	sa := option.WithCredentialsFile("path/to/serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return err
	}

	iter := client.Collection("breads").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if doc.Data()["Id"] == saveData.Id {
			return fmt.Errorf("既に保存されているデータです。")
		}
	}

	_, _, err = client.Collection("breads").Add(ctx, saveData)
	if err != nil {
		return fmt.Errorf("データの保存に失敗しました。: %v", err)
	}

	defer client.Close()
	return nil
}
