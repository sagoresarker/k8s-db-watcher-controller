-- Create the notification table
CREATE TABLE IF NOT EXISTS image_notifications (
    id SERIAL PRIMARY KEY,
    image_name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the notification channel
CREATE OR REPLACE FUNCTION notify_image_change() RETURNS TRIGGER AS $$
BEGIN
    PERFORM pg_notify('image_notify_channel', NEW.image_name);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create the trigger
CREATE TRIGGER image_notification_trigger
AFTER INSERT ON image_notifications
FOR EACH ROW EXECUTE FUNCTION notify_image_change();