CREATE TABLE "public"."user_vocab_bank" (
    "id" int2 NOT NULL,
    "value" varchar,
    "word_class" varchar,
    "meaning" varchar,
    "example" text,
    "explanation" text,
    "is_learned_status" bool,
    "user_id" uuid NOT NULL,
    "ipa" varchar NOT NULL,
    "created_at" date,
    "updated_at" date,
    CONSTRAINT "vocab_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id")
);