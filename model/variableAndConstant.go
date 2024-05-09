package model

type QUEUE string

const (
	PRODUCT_TO_ELASTIC        QUEUE = "product_to_elastic"
	UPDATE_PRODUCT_TO_ELASTIC QUEUE = "update_product_to_elastic"
	DELETE_PRODUCT_TO_ELASTIC QUEUE = "delete_product_to_elastic"
)

type INDEX_ELASTIC string

const (
	PRODUCT_INDEX INDEX_ELASTIC = "product_index"
)
