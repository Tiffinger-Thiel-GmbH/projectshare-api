import React, { useEffect, useState } from 'react';
import ProjectTable from '../components/ProjectTable';
import DocumentTable from '../components/DocumentTable';
import { Grid, Container, Box } from '@material-ui/core';
import { DocumentDTO, getDocumentsByProjectId, getProjects, ProjectDTO } from '../api/apis';

export default function Content() {
  const [isDrawerOpen, setIsDrawerOpen] = useState(false);
  const [selectedProject, setSelectedProject] = useState<string>();
  const [projects, setProjects] = useState<ProjectDTO[]>([]);
  const [documents, setDocuments] = useState<DocumentDTO[]>([]);

  useEffect(() => {
    getProjects().then(setProjects);
  }, []);

  useEffect(() => {
    if (selectedProject) getDocumentsByProjectId(selectedProject).then(setDocuments);
  }, [selectedProject]);

  return (
    <>
      <Container maxWidth="xl">
        <Grid container spacing={4} alignItems="center">
          <Grid item xs={6}>
            <Box ml={8}>
              {projects && !!projects.length && (
                <ProjectTable setIsDrawerOpen={setIsDrawerOpen} setSelectedProject={setSelectedProject} projects={projects} />
              )}
            </Box>
          </Grid>
          <Grid item xs={6}>
            <Box mr={8}>{documents && !!documents.length && <DocumentTable selectedProject={selectedProject} documents={documents} />}</Box>
          </Grid>
        </Grid>
      </Container>
    </>
  );
}
