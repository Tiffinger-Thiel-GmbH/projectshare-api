import React, { useState } from 'react';
import ProjectTable from '../components/ProjectTable';
import DocumentTable from '../components/DocumentTable';
import { Grid, Container, Box } from '@material-ui/core';

export default function Content() {
  const [isDrawerOpen, setIsDrawerOpen] = useState(false);
  const [selectedProject, setSelectedProject] = useState<number>();
  return (
    <>
      <Container maxWidth="xl">
        <Grid container spacing={4} alignItems="center">
          <Grid item xs={6}>
            <Box ml={8}>
              <ProjectTable setIsDrawerOpen={setIsDrawerOpen} setSelectedProject={setSelectedProject} />
            </Box>
          </Grid>
          <Grid item xs={6}>
            <Box mr={8}>{isDrawerOpen && selectedProject !== undefined && <DocumentTable selectedProject={selectedProject} />}</Box>
          </Grid>
        </Grid>
      </Container>
    </>
  );
}
