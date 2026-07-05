-- Caravan/Ryoudama points, PvP-area points, and Caravan Skill per character.
-- Field names and the points/pvp_points split come from mhfo-hd.dll RE
-- (see .claude/agent-memory/erupe-maintainer/project_caravan_packets.md):
-- hud_draw_caravan_points_label reads a normal points global, substituting a
-- separate value when the player is in a PvP area; Dialog_GetTokenSubstitution_
-- CaravanSkillName indexes a per-character "caravan skill" byte.
CREATE TABLE IF NOT EXISTS caravan (
    char_id    INTEGER PRIMARY KEY REFERENCES characters(id) ON DELETE CASCADE,
    points     INTEGER NOT NULL DEFAULT 0,
    pvp_points INTEGER NOT NULL DEFAULT 0,
    skill_id   SMALLINT NOT NULL DEFAULT 0
);

-- Guild-aggregate ("Ryoudan"/team) caravan points for the team ranking leaderboard,
-- mirroring the existing guilds.tower_rp column added in 0002_catch_up_patches.sql.
ALTER TABLE IF EXISTS guilds ADD COLUMN IF NOT EXISTS ryoudan_points INTEGER NOT NULL DEFAULT 0;
