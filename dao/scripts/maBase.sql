


			CREATE TABLE IF NOT EXISTS tenants (
				id_tenant SERIAL PRIMARY KEY,
				nom_tenant VARCHAR(50) NOT NULL
			

			);


			CREATE TABLE IF NOT EXISTS bibliotheques (
				id_bibliotheque SERIAL PRIMARY KEY,
				nom_bibliotheque VARCHAR(50) NOT NULL,
				id_tenant INT REFERENCES tenants(id_tenant)

			);
	
			CREATE TABLE IF NOT EXISTS livres (
				id_livre SERIAL PRIMARY KEY,
				auteur_livre	 VARCHAR(50) NOT NULL,
				titre_livre VARCHAR(50) NOT NULL,
				desc_livre VARCHAR(100) NOT NULL,
				id_tenant INT REFERENCES tenants(id_tenant)

			);

		CREATE TABLE IF NOT EXISTS revues (
				id_revue SERIAL PRIMARY KEY,
				auteur_revue VARCHAR(50) NOT NULL,
				titre_revue VARCHAR(50) NOT NULL,
				desc_revue VARCHAR(100) NOT NULL,
				id_tenant INT REFERENCES tenants(id_tenant)

			);
			CREATE TABLE IF NOT EXISTS magazines (
				id_magazine SERIAL PRIMARY KEY,
				auteur_magazine VARCHAR(50) NOT NULL,
				titre_magazine VARCHAR(50) NOT NULL,
				desc_magazine VARCHAR(100) NOT NULL,
				id_tenant INT REFERENCES tenants(id_tenant)

			);

		CREATE TABLE IF NOT EXISTS emprunts (
				id_emprunt			 SERIAL PRIMARY KEY,
				Date_emprunt    	DATE,
				id_tenant INT REFERENCES tenants(id_tenant)

			);

			CREATE TABLE IF NOT EXISTS clients (
				id_client SERIAL PRIMARY KEY,
				nom_client VARCHAR(50) NOT NULL,
				email_client 	VARCHAR(50) NOT NULL,
				password_user 	VARCHAR(100) NOT NULL,
				id_tenant INT REFERENCES tenants(id_tenant)
			);


			CREATE TABLE IF NOT EXISTS users (
				id_user SERIAL PRIMARY KEY,
				nom_user VARCHAR(50) NOT NULL,
				email_user	VARCHAR(50) NOT NULL,
				password_user VARCHAR(100) NOT NULL,
				id_role	INT REFERENCES 	roles(id_role),
				id_tenant INT REFERENCES tenants(id_tenant),
				id_session INT REFERENCES sessions(id_session)

			);


			
			CREATE TABLE IF NOT EXISTS roles (
				id_role SERIAL PRIMARY KEY,
				nom_role VARCHAR(50) NOT NULL,
				id_tenant INT REFERENCES tenants(id_tenant)
			);

			
			CREATE TABLE IF NOT EXISTS permission (
				id_permission SERIAL PRIMARY KEY,
				nom_permission VARCHAR(50) NOT NULL,
				id_tenant INT REFERENCES tenants(id_tenant)
			);

			
			CREATE TABLE IF NOT EXISTS roles_permissions (
				id_role_permission SERIAL PRIMARY KEY,
				id_role INT REFERENCES roles(id_role),
				id_permission INT REFERENCES permission(id_permission),
				id_tenant INT REFERENCES tenants(id_tenant)
			);
	
	
			CREATE TABLE IF NOT EXISTS sessions (
				id_session SERIAL PRIMARY KEY,
				nom_session VARCHAR(50) NOT NULL,
				start_time TIMESTAMP,
				duree		INTERVAL
			);

		
	