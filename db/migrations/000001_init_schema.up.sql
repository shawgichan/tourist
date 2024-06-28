CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "email" varchar(255) NOT NULL,
  "username" varchar(255) NOT NULL,
  "hashed_password" varchar(255) NOT NULL,
  "status" bigint NOT NULL,
  "roles_id" bigint NOT NULL,
  "profiles_id" bigint NOT NULL,
  "user_types_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "profiles" (
  "id" BIGSERIAL PRIMARY KEY,
  "first_name" varchar(255) NOT NULL,
  "last_name" varchar(255) NOT NULL,
  "addresses_id" bigint NOT NULL,
  "profile_image_url" varchar(255) NOT NULL,
  "phone_number" varchar(255) NOT NULL,
  "company_number" varchar(255) NOT NULL,
  "whatsapp_number" varchar(255) NOT NULL,
  "gender" bigint NOT NULL,
  "all_languages_id" bigint[] NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "ref_no" varchar NOT NULL,
  "cover_image_url" varchar
);

CREATE TABLE "locations" (
  "id" BIGSERIAL PRIMARY KEY,
  "lat" varchar(255) NOT NULL DEFAULT '0.0',
  "lng" varchar(255) NOT NULL DEFAULT '0.0',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_preferences" (
  "id" BIGSERIAL PRIMARY KEY,
  "users_id" bigint,
  "travel_companion_id" smallint,
  "place_category_preference" varchar,
  "price_range" varchar,
  "event_type_preference" varchar,
  "interests" text,
  "budget" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "places" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "description" text,
  "opening_hours" timestamptz NOT NULL DEFAULT (now()),
  "closing_hours" timestamptz NOT NULL DEFAULT (now()),
  "rating" decimal(2,1),
  "ticket_category" smallint NOT NULL,
  "ticket_price" varchar NOT NULL,
  "location_id" bigint NOT NULL,
  "place_type_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "cover_image_url" varchar(255) NOT NULL,
  "profile_image_url" varchar(255) NOT NULL,
  "resturant_branch_id" bigint NOT NULL,
  "preference_match" bigint[]
);

CREATE TABLE "restaurant_branches" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "resturant_id" bigint
);

CREATE TABLE "restaurants" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "stop_off_points" (
  "id" BIGSERIAL PRIMARY KEY,
  "trip_id" bigint,
  "places_id" bigint,
  "sequence_number" smallint,
  "arrival_time" timestamptz,
  "departure_time" timestamptz,
  "notes" text,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "trips" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" bigint,
  "name" varchar(255) NOT NULL,
  "description" text,
  "origin_location_id" bigint,
  "destination_location_id" bigint,
  "primary_place_id" bigint,
  "start_date" date NOT NULL,
  "end_date" date,
  "status" bigint,
  "travel_mode" bigint,
  "estimated_cost" decimal(10,2),
  "visibility" bigint,
  "shared_with" text[],
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "trip_feedback" (
  "id" BIGSERIAL PRIMARY KEY,
  "trip_id" bigint,
  "user_id" bigint,
  "rating" decimal(2,1),
  "feedback" text,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "places_reviews" (
  "id" BIGSERIAL PRIMARY KEY,
  "place_id" bigint,
  "user_id" bigint,
  "rating" decimal(2,1),
  "review" text,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "events_reviews" (
  "id" BIGSERIAL PRIMARY KEY,
  "place_id" bigint,
  "user_id" bigint,
  "rating" decimal(2,1),
  "review" text,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "events" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "description" text,
  "start_date" date NOT NULL,
  "end_date" date,
  "places_id" bigint,
  "event_type" bigint,
  "organizer_name" varchar(255),
  "organizer_website" varchar(255),
  "ticket_url" varchar(255),
  "image_id" bigint,
  "is_free" bool DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "places_events" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar,
  "places_id" bigint,
  "date" timestamptz,
  "event_type" bigint,
  "event_id" bigint
);

CREATE TABLE "event_trip" (
  "id" BIGSERIAL PRIMARY KEY,
  "event_id" bigint,
  "trip_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_following" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" bigint,
  "place_id" bigint
);

CREATE UNIQUE INDEX "uc_users_email" ON "users" ("email");

CREATE UNIQUE INDEX "uc_users_username" ON "users" ("username");

CREATE UNIQUE INDEX "uc_profiles_ref_no" ON "profiles" ("ref_no");

CREATE UNIQUE INDEX "uc_user_preferences_users_is" ON "user_preferences" ("users_id");

CREATE UNIQUE INDEX ON "user_following" ("user_id", "place_id");

ALTER TABLE "trips" ADD FOREIGN KEY ("primary_place_id") REFERENCES "places" ("id");

ALTER TABLE "event_trip" ADD FOREIGN KEY ("id") REFERENCES "events" ("id");

ALTER TABLE "event_trip" ADD FOREIGN KEY ("trip_id") REFERENCES "trips" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("places_id") REFERENCES "places" ("id");

ALTER TABLE "places_reviews" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "places_reviews" ADD FOREIGN KEY ("place_id") REFERENCES "places" ("id");

ALTER TABLE "events_reviews" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "events_reviews" ADD FOREIGN KEY ("place_id") REFERENCES "events" ("id");

ALTER TABLE "trip_feedback" ADD FOREIGN KEY ("trip_id") REFERENCES "trips" ("id");

ALTER TABLE "trip_feedback" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "trips" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "stop_off_points" ADD FOREIGN KEY ("trip_id") REFERENCES "trips" ("id");

ALTER TABLE "trips" ADD FOREIGN KEY ("origin_location_id") REFERENCES "locations" ("id");

ALTER TABLE "trips" ADD FOREIGN KEY ("destination_location_id") REFERENCES "locations" ("id");

ALTER TABLE "stop_off_points" ADD FOREIGN KEY ("places_id") REFERENCES "places" ("id");

ALTER TABLE "places" ADD FOREIGN KEY ("location_id") REFERENCES "locations" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("profiles_id") REFERENCES "profiles" ("id");

ALTER TABLE "user_preferences" ADD FOREIGN KEY ("users_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "places" ADD FOREIGN KEY ("resturant_branch_id") REFERENCES "restaurant_branches" ("id") ON DELETE CASCADE;

ALTER TABLE "restaurant_branches" ADD FOREIGN KEY ("resturant_id") REFERENCES "restaurants" ("id") ON DELETE CASCADE;

ALTER TABLE "places_events" ADD FOREIGN KEY ("places_id") REFERENCES "places" ("id") ON DELETE CASCADE;

ALTER TABLE "places_events" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON DELETE SET NULL;

ALTER TABLE "user_following" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "user_following" ADD FOREIGN KEY ("place_id") REFERENCES "places" ("id") ON DELETE CASCADE;
