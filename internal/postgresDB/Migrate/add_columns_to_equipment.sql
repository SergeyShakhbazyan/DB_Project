ALTER TABLE "MyDatabase".public.equipment
    ADD COLUMN warranty_period INTEGER,
    ADD COLUMN last_maintenance_date DATE;

ALTER TABLE "MyDatabase".public.equipment
    DROP COLUMN warranty_period,
    DROP COLUMN last_maintenance_date;